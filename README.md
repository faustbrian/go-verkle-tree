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

## Installation

```sh
go get github.com/faustbrian/go-verkle-tree
```

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

Use the [documentation index](docs/README.md), [profile freeze](specification/profile-freeze.md),
[backend audit](docs/backend-audit.md), and [threat model](docs/threat-model.md)
before adoption. The [detailed reference](docs/reference.md) preserves the full
proof, storage, witness, recovery, and profile contracts.

## Development

Run `make check` and the package conformance gates before changing profile,
commitment, proof, encoding, or storage behavior.

## License

MIT. See [LICENSE](LICENSE).
