# Security Policy

## Experimental Status And Known Limitations

**This library is experimental and is not production-ready. Do not use it to
protect production integrity or confidentiality.** Stable Go API versions and
published Git tags do not certify cryptographic soundness, side-channel safety,
or production suitability.

Backend remediation is deferred while this repository is used for research.
The maintainers own the following unresolved risks; documenting them is not a
fix or an assurance that applications can safely work around them:

| Known limitation | Consequence and current containment | Future remediation |
| --- | --- | --- |
| The pinned IPA backend lacks an independent production audit. | Complete proof soundness is not established. Restrict use to isolated, non-production experiments without security-sensitive assets. | Adopt a maintained, independently reviewed backend compatible with a fully specified profile; do not implement replacement commitment arithmetic in this tree library. |
| Degenerate transcript challenges have unresolved reject-or-retry semantics. | Exceptional challenge behavior is not proven sound or interoperable. No practical challenge preimage is demonstrated by this audit; this remains an unresolved cryptographic boundary. | Specify a versioned failure or retry rule and independently reproduce positive and exceptional-challenge behavior. |
| An admitted dependency proof call cannot be cancelled and dependency work may derive parallelism from CPU count. | A caller deadline does not interrupt all in-flight CPU work. Use small explicit budgets and isolated experiments; a timeout wrapper does not cancel the underlying work. | Require backend-level cancellation and explicit worker/resource budgets, with lifecycle and cancellation regression tests. |
| Exported mutable dependency globals remain authoritative. | The backend does not provide a fully immutable per-instance cryptographic configuration. Do not mutate dependency globals or share experimental execution with untrusted application code. | Use an immutable, instance-owned backend configuration and verify concurrent isolation. |
| Complete side-channel and native-platform behavior is unverified. | No constant-time or cross-platform production security claim is made. Do not process production secrets or treat compile-only platform checks as runtime assurance. | Obtain a scoped side-channel audit and native execution evidence for every supported production target. |

The [backend audit](docs/backend-audit.md), [threat model](docs/threat-model.md),
and [platform audit](docs/platforms.md) contain the supporting boundaries.
Revisit this deferral before any production adoption or production-readiness
claim, and whenever the backend, transcript profile, or relevant dependency
changes. Production qualification requires resolving these limitations,
independent cryptographic review, and verification of the affected public
contracts and consumers; passing ordinary CI alone is insufficient.

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

The latest stable `v1` release receives fixes within this experimental scope;
no version is supported for production use. Reports affecting `main` are also
welcome, but `main` is not a supported release. The compatibility
and deprecation boundaries are documented in
[`COMPATIBILITY.md`](COMPATIBILITY.md) and [`DEPRECATION.md`](DEPRECATION.md).

API compatibility does not override the experimental status or the known
limitations above.

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
