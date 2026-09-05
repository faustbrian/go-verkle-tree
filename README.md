# verkle-tree

[![CI](https://github.com/faustbrian/go-verkle-tree/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/faustbrian/go-verkle-tree/actions/workflows/ci.yml)
[![CodeQL](https://img.shields.io/badge/CodeQL-required-blue)](https://github.com/faustbrian/go-verkle-tree/actions/workflows/ci.yml)
[![Coverage](https://img.shields.io/badge/coverage-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Mutation](https://img.shields.io/badge/mutation-100%25_required-blue)](CONTRIBUTING.md#verification)
[![Documentation](https://img.shields.io/badge/docs-checked_in_CI-blue)](docs/)
[![Go Reference](https://pkg.go.dev/badge/github.com/faustbrian/go-verkle-tree.svg)](https://pkg.go.dev/github.com/faustbrian/go-verkle-tree)
[![Release](https://img.shields.io/github/v/release/faustbrian/go-verkle-tree?sort=semver)](https://github.com/faustbrian/go-verkle-tree/releases)
[![Go](https://img.shields.io/badge/go-1.26.6-00ADD8?logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

`verkle-tree` is a storage-independent authenticated key/value tree backed by
vector commitments. Its v1 Go API exposes the package-owned research profile
`verkletree-bandersnatch-ipa-256-v0`; it does not claim compatibility with an
eventual Ethereum production Verkle profile.

The pinned cryptographic backend has not received the independent audit
required for production use. Treat the implementation and profile as
research-grade despite the stable Go API and explicit compatibility contract.

## Status, package, and platform

The module has a stable v1 Go API and requires Go 1.26.6. Its only public
package is the root package, whose default import identifier is `verkletree`.
All subpackages are internal implementation details.

Native runtime acceptance evidence is limited to `darwin/arm64`. Other native
targets in the [platform audit](docs/platforms.md) have compile-only evidence;
Rosetta results are diagnostic and do not establish native `darwin/amd64`
support. The module has no cgo path, but its cryptographic dependencies use
architecture-specific assembly. No constant-time or complete side-channel
claim is made.

## Installation

```sh
go get github.com/faustbrian/go-verkle-tree
```

## When to use it

Use `verkletree` when an application accepts the fixed research profile and
needs bounded immutable snapshots, profile-bound proofs and witnesses,
stateless updates, or caller-owned storage protocols. Do not use it as an
Ethereum compatibility layer, an audited production cryptography component,
or a storage adapter. The [adoption and migration guide](docs/adoption.md)
defines the decision and rollback boundaries.

## Quick start

The API has no unbounded defaults. Callers must select finite state, tree, and
commitment limits before constructing a snapshot:

```go
profile := verkletree.BandersnatchIPA256V0()
limits := verkletree.SnapshotLimits{
	State: verkletree.StateLimits{
		MaxEntries: 64, MaxBatchUpdates: 64,
		MaxTemporaryBytes: 16 << 20,
	},
	Tree: verkletree.TreeLimits{
		MaxEntries: 64, MaxStems: 64, MaxNodes: 128, MaxEdges: 128,
		MaxCommitments: 256, MaxFieldMappings: 256,
		MaxCommitmentTerms: 1 << 16, MaxTemporaryBytes: 16 << 20,
	},
	Commitment: verkletree.CommitmentLimits{
		MaxGeneratorDerivations: 256, MaxScalarDecodes: 256,
		MaxMSMTerms: 256, MaxTemporaryBytes: 1 << 20,
	},
}

var key verkletree.Key
snapshot, err := verkletree.NewSnapshot(
	ctx,
	profile,
	[]verkletree.Entry{{Key: key, Value: verkletree.Value{}}},
	limits,
)
if err != nil {
	return err
}
```

The checked-in [`ExampleSnapshot`](example_test.go) is the executable version
of this flow and is compiled and run by the Go example test gate.

## Operational contract

- Construction selects the fixed `BandersnatchIPA256V0` profile and validates
  explicit limits before attacker-amplified work. See the
  [usage guide](docs/usage.md) and [detailed reference](docs/reference.md).
- Expensive public operations accept a context. Queued work observes
  cancellation, but an admitted dependency proof call cannot be interrupted.
- Sentinel errors work with `errors.Is`; `ResourceError` and
  `StoreCapabilityError` work with `errors.As`.
- Immutable snapshots, engines, proofs, witnesses, and successful results are
  safe for concurrent use. Callers construct, retain, and discard engines;
  they expose no close lifecycle and start no background work outside an
  operation. Callers own storage implementations, read-view closure, writer
  coordination, durability, publication, and recovery.
- The repository provides no concrete storage adapter and declares no owned
  Golib consumer. Applications integrate through the public caller-owned
  storage interfaces and must validate those implementations independently.

## Guarantees and limitations

- Snapshots and transitions are immutable and profile-bound.
- Set, delete, encoding, reconstruction, proof, storage, recovery, retention,
  and pruning work is explicitly bounded.
- Caller-owned storage writes are atomic and authenticated against published
  roots.
- Membership and non-membership proofs bind the selected package profile.
- The implementation rebuilds complete tree state for maintained stateful
  updates and does not provide concrete storage adapters.
- The package cannot restore missing or corrupt published state.

## Documentation

Use the [documentation index](docs/README.md) for the complete guide set. Start
with [adoption and FAQ](docs/adoption.md), the [usage guide](docs/usage.md),
the [detailed reference](docs/reference.md), and the
[executable example](example_test.go). Review [storage operations](docs/storage-operations.md),
[platform evidence](docs/platforms.md), [compatibility](docs/compatibility.md),
[performance](docs/benchmarks.md), and the [threat model](docs/threat-model.md)
before integration. Project routes include [support](SUPPORT.md),
[private security reporting](SECURITY.md), the [changelog](CHANGELOG.md), and
the [license](LICENSE).

The [specification decision register](docs/specification-decisions.md),
[profile freeze](specification/profile-freeze.md), and
[backend audit](docs/backend-audit.md) define the research-profile boundary.

For ecosystem-wide package selection and ownership conventions, see the
[versioned Golib ecosystem index](https://github.com/faustbrian/go-library-tools/blob/v1.4.0/docs/ecosystem/README.md)
and its [Domain utilities family](https://github.com/faustbrian/go-library-tools/blob/v1.4.0/docs/ecosystem/design-language.md#package-families-and-selection).

## Development

Run `make check` and the package conformance gates before changing profile,
commitment, proof, encoding, or storage behavior.

Run `make cohesion` to validate the module's family, ownership, lifecycle,
documentation, and supported-environment metadata.

## License

MIT. See [LICENSE](LICENSE).
