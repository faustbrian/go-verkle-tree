# Security Policy

## Reporting

Do not open a public issue for a suspected vulnerability. Report it privately
through [GitHub Security Advisories for
`faustbrian/go-verkle-tree`](https://github.com/faustbrian/go-verkle-tree/security/advisories/new)
before public disclosure. Do not include exploit details, credentials, private
fixtures, or affected deployment information in a public report.

Include the affected module and version, impact, reproduction, preconditions,
and any suggested mitigation. Reports are acknowledged as soon as practical;
timelines depend on severity and verification.

## Supported Versions

The latest stable `v1` release receives security fixes. Older releases and the
`main` branch are not supported; upgrade before reporting. The compatibility
and deprecation boundaries are documented in
[`COMPATIBILITY.md`](COMPATIBILITY.md) and [`DEPRECATION.md`](DEPRECATION.md).

The pinned cryptographic backend is not independently audited, and complete
side-channel behavior is unverified. See the [threat model](docs/threat-model.md)
and [platform audit](docs/platforms.md) before assessing exposure.

## Security Gates

Releases require isolated tests, race and hostile-input checks, exact coverage
and mutation results, `govulncheck`, secret scanning, license verification,
SBOM generation, provenance validation, and clean-consumer resolution. A
missing scanner or unavailable service is a failed gate, not a warning.

Security fixes MUST include a regression test that does not publish weaponized
details or real secrets. Credentials MUST be redacted from logs and evidence.

## Repository Assurance

The repository [safety and concurrency policy](AGENTS.md#safety-and-concurrency)
and [supply-chain policy](AGENTS.md#dependencies-and-supply-chain) define shared
trust boundaries and release requirements. Package-specific security guidance
refines those rules for its owned boundary.
