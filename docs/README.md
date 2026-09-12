# Documentation

`go-verkle-tree` exposes one stable-v1 public package, imported with the
default identifier `verkletree`, and requires Go 1.27.0. It implements only
the fixed `verkletree-bandersnatch-ipa-256-v0` research profile. Native runtime
acceptance evidence is limited to `darwin/arm64`; other native targets are
compile-only. It provides no concrete storage adapter and makes no Ethereum,
production-suitability, constant-time, or complete side-channel claim.

## Getting started

- [Adoption](adoption.md)
- [Usage](usage.md)
- [Executable example](../example_test.go)
- [FAQ](adoption.md#faq)

## API and construction

- [Detailed package reference](reference.md)
- [API boundaries](api-boundaries.md)
- [API audit](api-audit.md)
- [Research profile freeze](../specification/profile-freeze.md)
- [Specification decisions](specification-decisions.md)
- [Specification and conformance matrix](../specification/README.md)

## Concepts and design

- [Backend audit](backend-audit.md)
- [Platform and CPU evidence](platforms.md)

## Operations and security

- [Storage operations](storage-operations.md)
- [Threat model](threat-model.md)
- [Security reporting](../SECURITY.md)
- [Support](../SUPPORT.md)

## Compatibility and performance

- [Benchmarks](benchmarks.md)
- [Compatibility](compatibility.md)
- [Compatibility policy](../COMPATIBILITY.md)
- [Deprecation policy](../DEPRECATION.md)

## Project and verification

- [Contribution guide](../CONTRIBUTING.md)
- [Development and verification](development.md)
- [Release history](../CHANGELOG.md)
- [License](../LICENSE)
