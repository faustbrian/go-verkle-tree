# Contributing

## Before Editing

1. Read [`AGENTS.md`](AGENTS.md), [SECURITY.md](SECURITY.md), and affected docs.
2. Run focused standalone Go tests for the affected package.
3. Identify owned dependencies and reverse dependants in `modules.json`.
4. Preserve unrelated work and generated/corpus provenance.

## Changes

Keep commits focused and conventional. Update every affected changelog with
the behavior and migration impact. Public API changes require compatibility
evidence and documentation. Specification behavior requires a decision record,
fixture coverage, and interoperability evidence.

New direct dependencies and dependency updates must follow the
[dependency governance policy](AGENTS.md#dependencies-and-supply-chain). Package-local
update bots are forbidden; the root policy owns every module and action update.

Specification-backed changes must follow the
[specification governance contract](AGENTS.md#design), update
the affected stable decision entries, and complete the Specification Decisions
section of the pull request template. An unresolved interpretation or stale
source pin is release-blocking; peer behavior cannot silently select policy.
The root module's decisions and conformance boundary are recorded in the
[specification decision register](docs/specification-decisions.md).

Explicit full mutation qualification must finish with zero surviving viable
mutants; it is not required for every ordinary process or documentation edit.

Do not fork shared verification workflows, add permanent replacements, machine-specific
paths, bypass flags, broad mutation exclusions, or aggregate quality metrics
that hide a failing package.

## Verification

This library remains experimental and unsuitable for production integrity or
confidentiality. Ordinary verification, stable APIs and published versions do
not resolve the accepted backend, transcript, cancellation, mutable-global or
side-channel limitations in [SECURITY.md](SECURITY.md).

Use Go 1.27.2 for verification; the public module minimum remains Go 1.27.0.
Standalone tests need neither sibling repositories nor building the verifier
release from source:

```bash
GOWORK=off go test ./... -count=1 -timeout=3m
```

`make check` and `make ci` use the checksum-qualified public CLI and bounded
`--local --all` verification: deterministic tests, saved fuzz seeds and the
named local risk check, without active fuzz exploration. Committed-range CI
adds `--base-sha` to select active smoke; the explicit single-target command
below also explores. No services, test tags or goal ledger are prerequisites.
The existing hosted 386 storage-width job remains jointly required with shared
CI; it checks accounting/update conversions, not cryptographic soundness or
native production support.

The selected cheap state-input regression runs once within two minutes without
race instrumentation or a cryptographic backend campaign. It checks batch,
scratch and result-budget refusals, including the exact resource kind and
temporary-byte accounting:

```bash
GOWORK=off go test ./internal/statemodel -run '^TestSnapshotEnforcesOperationBudgetsBeforeAllocation$' -count=1 -timeout=2m
```

Normal tests run saved fuzz seeds, not active exploration. For a changed
state-transition input boundary, select one target; optional local mutation should
stay in one affected package:

```bash
GOWORK=off go test ./internal/statemodel -run '^$' -fuzz '^FuzzSnapshotTransitions$' -fuzztime=2s -timeout=2m
golib mutation --module . --package ./internal/statemodel
```

Agent-run Go commands must use and remove a task-owned disposable `GOCACHE`.
Shared PR, main and scheduled CI are bounded. Explicit `make check-full` and
`release_dry_run` retain all original full operations: 23 fuzz targets at
10,000 iterations/20-minute limits, full mutation/coverage/race/API/security,
20-minute internal conformance, 30-minute pinned Go/Rust interoperability,
and the 100ms/20-minute benchmark operation. No full cryptographic mutation
campaign is required merely to deliver this CI-policy adoption.

Interoperability needs the pinned Rust toolchain and Cargo sources, Go, jq,
Git, diff and shasum; reference and fixture checksums/refusals remain intact.
`make specification-check` is offline; `make specification-check-online` and
full qualification retain online authority. Use module-selected
`golib release check --module .` and `golib release dry-run --module .` only at
the deliberate release boundary; rehearsal does not publish or qualify the
research backend for production. See [docs/development.md](docs/development.md).
Report unavailable or failed selected checks without implying production
security or release readiness from the ordinary route.

## Selected Monthly Assurance

The owner-level [selected-assurance caller](.github/workflows/assurance.yml) is
configured for main only, on calendar day 29 when present, at 08:17 UTC. It selects one
of the 23 declared root-module package-target identities using the shared
UTC-week selection, with 30 seconds of active exploration. Monthly selection
samples one identity; it does not promise exhaustive rotation. The selected
risks are state-transition/reference equivalence, canonical snapshot/proof
encodings, bounded hostile input and storage/layout histories. The two package-owned
`FuzzSnapshotTransitions` identities remain distinct. This is an isolated,
non-production diagnostic, not cryptographic certification or full research
qualification. All backend, transcript, non-cancellable dependency work,
mutable-global and side-channel limitations in [SECURITY.md](SECURITY.md) remain
unresolved and unchanged.

Base-free ordinary checks execute deterministic seeds, not active exploration;
committed-range CI selects affected active smoke. Manual assurance requires a
declared selector from the root-module fuzz operation in `.golib.yaml`, for
example `./internal/statemodel::^FuzzSnapshotTransitions$`. A bare anchored
target is accepted only when unique across declared packages; ambiguous, empty, unknown
and malformed selections refuse. The shared job bounds setup and execution
to ten minutes, active workers and CPU to two, owned memory to 2 GiB, native
output to 1 MiB and private-cache lifetime. Failure, cancellation, exhausted
budgets and cleanup refusal remain attributable failures. This lane does not
run mutation or replace Required, original full fuzz/race budgets,
explicit `make check-full`, or relevant public-consumer/release qualification.

## Adding A Module

Follow [repository structure policy](AGENTS.md#repository-structure). New modules
require an explicit purpose, ownership boundary, dependency review, package
catalog entry, full quality gates, documentation, changelog, license, security
policy, compatibility plan, and release dry-run.
