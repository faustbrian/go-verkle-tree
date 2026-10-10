package verkletree_test

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"runtime"
	"slices"
	"strings"
	"testing"
)

func TestProductionSourceHasNoPackageInitOrOwnedGoroutines(t *testing.T) {
	t.Parallel()

	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("resolve architecture test source path")
	}
	moduleRoot := filepath.Dir(filename)
	violations, err := scanProductionSourceArchitecture(moduleRoot)
	if err != nil {
		t.Fatalf("scan production source: %v", err)
	}
	for _, violation := range violations {
		t.Errorf(
			"package-owned production source must not %s: %s",
			violation.rule,
			violation.position,
		)
	}
}

func TestProductionSourceArchitectureRespectsModuleOwnership(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	for name, content := range map[string]string{
		"go.mod":                    "module fixture.test/root\n",
		"owned.go":                  "package fixture\nfunc init() {}\n",
		"internal/owned.go":         "package internal\nfunc launch() { go func() {}() }\n",
		".golib-tooling/go.mod":     "module fixture.test/tooling\n",
		".golib-tooling/foreign.go": "package tooling\nfunc init() {}\nfunc launch() { go func() {}() }\n",
	} {
		path := filepath.Join(root, filepath.FromSlash(name))
		if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
			t.Fatal(err)
		}
	}
	violations, err := scanProductionSourceArchitecture(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(violations) != 2 {
		t.Fatalf("owned violations = %d, want 2", len(violations))
	}
	for _, violation := range violations {
		if strings.Contains(violation.position.Filename, ".golib-tooling") {
			t.Fatal("independent module classified as owned source")
		}
	}
}

func scanProductionSourceArchitecture(moduleRoot string) ([]sourceArchitectureViolation, error) {
	violations := make([]sourceArchitectureViolation, 0)
	err := filepath.WalkDir(moduleRoot, func(
		path string,
		entry fs.DirEntry,
		walkErr error,
	) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() && path != moduleRoot {
			// A nested module owns its source independently, including the CI
			// tooling checkout. Do not apply this module's architecture to it.
			if _, err := os.Stat(filepath.Join(path, "go.mod")); err == nil {
				return filepath.SkipDir
			} else if !os.IsNotExist(err) {
				return err
			}
		}
		if entry.IsDir() ||
			filepath.Ext(path) != ".go" ||
			strings.HasSuffix(path, "_test.go") {
			return nil
		}

		fileSet := token.NewFileSet()
		file, err := parser.ParseFile(fileSet, path, nil, parser.SkipObjectResolution)
		if err != nil {
			return fmt.Errorf("parse production source %s: %w", path, err)
		}
		violations = append(
			violations,
			findSourceArchitectureViolations(fileSet, file)...,
		)

		return nil
	})
	return violations, err
}

func TestSourceArchitectureViolationDetection(t *testing.T) {
	t.Parallel()

	const source = `package fixture

func init() {}

func launch() {
	go func() {}()
}
`
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(
		fileSet,
		"representative_violation.go",
		source,
		parser.SkipObjectResolution,
	)
	if err != nil {
		t.Fatalf("parse representative violation: %v", err)
	}

	violations := findSourceArchitectureViolations(fileSet, file)
	rules := make([]string, len(violations))
	for index := range violations {
		rules[index] = violations[index].rule
	}
	slices.Sort(rules)
	want := []string{"define package init functions", "start goroutines"}
	if !slices.Equal(rules, want) {
		t.Fatalf("detected rules = %q, want %q", rules, want)
	}
}

type sourceArchitectureViolation struct {
	rule     string
	position token.Position
}

func findSourceArchitectureViolations(
	fileSet *token.FileSet,
	file *ast.File,
) []sourceArchitectureViolation {
	violations := make([]sourceArchitectureViolation, 0)
	ast.Inspect(file, func(node ast.Node) bool {
		switch value := node.(type) {
		case *ast.FuncDecl:
			if value.Recv == nil && value.Name.Name == "init" {
				violations = append(violations, sourceArchitectureViolation{
					rule:     "define package init functions",
					position: fileSet.Position(value.Pos()),
				})
			}
		case *ast.GoStmt:
			violations = append(violations, sourceArchitectureViolation{
				rule:     "start goroutines",
				position: fileSet.Position(value.Pos()),
			})
		}

		return true
	})

	return violations
}
