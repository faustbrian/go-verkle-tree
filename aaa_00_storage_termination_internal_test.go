package verkletree

import (
	"bytes"
	"context"
	"os"
	"os/exec"
	"regexp"
	"testing"
	"time"

	"github.com/faustbrian/go-verkle-tree/internal/authstate"
)

// Keep these assertions serial and before the parallel storage/proof fixtures.
// A divergent arithmetic regression must fail in a reaped child, not leave an
// already-running fixture spinning after the test runner observes a failure.
func TestStorageAccountingAndUpdateConversionTerminate(t *testing.T) {
	for _, scenario := range []string{"capacity growth", "bounded append", "capped page budget", "delete followed by set"} {
		t.Run(scenario, func(t *testing.T) {
			if storageAccountingOwnedChild(t) {
				return
			}
			switch scenario {
			case "capacity growth":
				const maximum = 1<<31 - 1
				for _, test := range []struct {
					length, capacity, additional, maximum, want int
				}{
					{3, 3, 1, 20, 6},
					{0, 0, 5, 6, 6},
					{0, 0, 2, maximum, 2},
					{0, 0, 1<<30 + 1, maximum, maximum},
					{maximum - 1, maximum - 1, 2, maximum, maximum},
				} {
					got := storageAuditResultCapacity(test.length, test.capacity, test.additional, test.maximum)
					if got != test.want {
						t.Fatalf("bounded result capacity = %d, want %d", got, test.want)
					}
				}
			case "bounded append":
				nodes := []NodeID{{1}, {2}, {3}, {4}}
				got := appendStorageAuditNode(nodes, NodeID{5}, 6)
				if len(got) != 5 || cap(got) != 6 {
					t.Fatalf("bounded append = (len %d, cap %d), want (5, 6)", len(got), cap(got))
				}
				for index, node := range got {
					if node != (NodeID{byte(index + 1)}) {
						t.Fatalf("bounded append lost node %d", index)
					}
				}
			case "capped page budget":
				limits := StorageAuditLimits{MaxUnreachableNodes: 4, MaxTemporaryBytes: 8 * storageAuditNodeIDBytes}
				got, err := storageAuditPageLimit(limits, 0, 0, 3, 4, 4)
				if err != nil || got != 4 {
					t.Fatalf("capped page admission = (%d, %v), want (4, nil)", got, err)
				}
			case "delete followed by set":
				key, value := Key{2}, Value{3}
				got, err := toInternalWitnessUpdates(context.Background(), []Update{Delete(Key{1}), Set(key, value)})
				want := []authstate.Update{authstate.Delete(authstate.Key{1}), authstate.Set(authstate.Key(key), authstate.Value(value))}
				if err != nil || len(got) != len(want) || got[0] != want[0] || got[1] != want[1] {
					t.Fatalf("delete-followed-by-set conversion = (%v, %v), want %v", got, err, want)
				}
			}
		})
	}
}

func storageAccountingOwnedChild(t *testing.T) bool {
	t.Helper()
	const selector = "GOLIB_STORAGE_ACCOUNTING_CHILD"
	if os.Getenv(selector) == t.Name() {
		return false
	}
	executable, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(t.Context(), 2*time.Second)
	defer cancel()
	command := exec.CommandContext(ctx, executable, "-test.run=^"+regexp.QuoteMeta(t.Name())+"$", "-test.count=1")
	command.Env = append(os.Environ(), selector+"="+t.Name())
	var output bytes.Buffer
	command.Stdout, command.Stderr = &output, &output
	if err := command.Run(); err != nil {
		t.Fatalf("owned accounting assertion failed or exceeded deadline: %v; %s", err, output.String())
	}
	return true
}
