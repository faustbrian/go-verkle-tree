# Development and Verification

The repository uses the released `go-library-tools` contract for local and CI
verification. Install the checksum-qualified public `golib` version declared in
[`.golib.yaml`](../.golib.yaml), or provide that version through the project
tooling environment, then run:

```sh
GOWORK=off go test ./... -count=1 -timeout=3m
make check
make ci
```

`make check` runs bounded `--local --all` verification. `make ci` adds repository,
cohesion and offline specification checks. Ordinary PR, main and scheduled
shared CI are bounded too; Required also retains the existing 386 accounting
job. Missing selected prerequisites fail closed. A sibling checkout, goal ledger
or verifier release build is not required for standalone Go contribution.

The selected local input-budget regression is
`TestSnapshotEnforcesOperationBudgetsBeforeAllocation` in `./internal/statemodel`,
count one with a two-minute timeout. It asserts batch/scratch/result resource
refusals and exact accounting, without a cryptographic campaign or manufactured
race requirement.

Normal tests run saved fuzz seeds, not active exploration. Select one affected
target for a two-second run as shown in [CONTRIBUTING.md](../CONTRIBUTING.md);
package-local mutation is optional for a named risk, not a universal edit gate.
`make check-full` and the explicit release dry-run retain full mutation and all
23 configured 10,000-iteration fuzz targets with their 20-minute limits.
`make specification-check-online` retains explicit online authority.

Verkle-specific conformance and interoperability checks remain package-owned
operations. They are defined in [`verification/package.mk`](../verification/package.mk)
and invoke the pinned Go and Rust reference harnesses in
[`interoperability/`](../interoperability/). The reference source and fixture
checksums in [`specification/sources.json`](../specification/sources.json) are
part of those verification contracts.

Full conformance keeps its 20-minute limit; pinned Go/Rust interoperability
keeps 30 minutes and requires Cargo, `rustc`, `jq`, Git, `diff` and `shasum` in addition
to Go. The benchmark operation retains 100ms/20m. These checks concern the fixed
research profile; they do not resolve the accepted cryptographic limitations
or establish production suitability. [SECURITY.md](../SECURITY.md) remains the
authority for experimental containment and deferred remediation.

Mutation checkpoints and their migration record are retained under
[`../.verification/mutation/`](../.verification/mutation/). They are imported
by the strict configuration without rerunning content-identical campaigns.
