# Development and Verification

The repository uses the released `go-library-tools` contract for local and CI
verification. Install `golib` v1.0.7 or provide it through the project tooling
environment, then run:

```sh
make inventory
make check
make ci
```

`make check` runs the complete package gate. `make ci` adds repository
validation before running the package gates. Both commands use the same
configuration and thresholds as GitHub Actions; missing tools, evidence,
fixtures, or required operations fail closed.

Verkle-specific conformance and interoperability checks remain package-owned
operations. They are defined in [`verification/package.mk`](../verification/package.mk)
and invoke the pinned Go and Rust reference harnesses in
[`interoperability/`](../interoperability/). The reference source and fixture
checksums in [`specification/sources.json`](../specification/sources.json) are
part of those verification contracts.

Mutation checkpoints and their migration record are retained under
[`../.verification/mutation/`](../.verification/mutation/). They are imported
by the strict configuration without rerunning content-identical campaigns.
