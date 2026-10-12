# Verkle profile conformance matrix

The root module implements the package-owned
[`verkletree-bandersnatch-ipa-256-v0`](bandersnatch-ipa-256-v0.md) research
profile. The [specification decision register](../docs/specification-decisions.md)
defines the supported boundary. It does not claim production suitability,
external audit, generic Verkle compatibility, or Ethereum protocol compliance.

[`monitoring.json`](monitoring.json) pins the profile, the four declared
Ethereum research inputs, and their change authorities, including the
`rust-verkle` commit feed, for review every 30 days.
[`sources.json`](sources.json) retains fixture and implementation
provenance. [`peer-assessment.json`](peer-assessment.json) records why the
reproducible Go and Rust corpora are not described as maintained-peer
differential conformance.

The 2026-10-12 upstream review found `rust-verkle` master still at the fixture's
exact revision `e27b8b4edf1992b4afa636c2fc7983bcc27ddb88` and recorded tree.
Its source and license bytes still match the provenance record. The repository
is not archived but has had no push since 2024-10-25; this is dormancy, not
evidence of active maintenance or cryptographic remediation. The three Cargo
Git inputs remain immutable fixture sources owned by `sources.json`, while
registry dependencies receive Dependabot review. A replacement or upstream
change requires source/license review, corpus regeneration through
[`scripts/check-interoperability.sh`](../scripts/check-interoperability.sh),
and profile/compatibility review; its cost
is re-establishing the pinned generator and differential evidence, not merely
changing a dependency version. This review does not refresh the recorded Cargo
advisory assessment.

## Decision conformance

| Decision | Authority | Executable evidence | Differential result |
| --- | --- | --- | --- |
| VERKLE-DEC-001 | `verkle-profile-v0-source` | `TestProfileIsExactAndRejectsOtherValues`, `TestRustVerkleEncodingVectors`, `TestRustVerkleGeneratorSet`, `TestCommitmentEngineMatchesPinnedRustVectors`, `TestBuildMatchesPinnedRustTreeRoots`, `TestStatelessUpdaterMatchesPinnedRustRebuiltTransitions`, `FuzzDecodeTreeProof`, `FuzzStatelessUpdaterMatchesStatefulTransition` | Not assessed: no currently maintained independent implementation covers the complete profile. |
| VERKLE-DEC-002 | `ethereum-eip-6800-source`, `ethereum-eip-4762-source`, `ethereum-eip-7612-source`, `ethereum-eip-7748-source` | `TestProfileIsExactAndRejectsOtherValues` | Not assessed: the EIPs are explicitly unimplemented research inputs, not an interoperability target. |

Pinned corpus agreement remains bounded to the exact fixtures named in
[`conformance.json`](conformance.json). A source change requires review of the
affected decision, public contract, fixtures, compatibility statement, and
changelog before any behavior or compliance claim changes.
