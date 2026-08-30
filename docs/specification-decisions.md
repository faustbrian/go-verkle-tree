# Verkle Tree Specification Decisions

This register separates the implemented package-owned research profile from
unimplemented Ethereum research inputs. The authoritative machine records are
[`decisions.json`](../specification/decisions.json),
[`conformance.json`](../specification/conformance.json), and
[`monitoring.json`](../specification/monitoring.json).

## VERKLE-DEC-001: Package-owned research profile identity

| Field | Decision |
| --- | --- |
| Status and owner | `resolved`; `go-verkle-tree` maintainers |
| Classification and scope | `interoperability policy`; `normative` |
| Specification and version | `verkletree-bandersnatch-ipa-256-v0 package-owned research profile v0`; `verkletree-bandersnatch-ipa-256-v0 at 79ab045f13f1ae8214bbbe10f12f434e944b5eee` |
| Authority and section | `verkle-profile-v0-source`; [complete document](https://raw.githubusercontent.com/faustbrian/go-verkle-tree/79ab045f13f1ae8214bbbe10f12f434e944b5eee/specification/bandersnatch-ipa-256-v0.md); `Complete document`; `MUST` |
| Issue | No stable external specification selects one complete caller-composable Verkle tree profile with this package's curve, generator, transcript, encoding, tree, proof, witness, and state-transition semantics. |
| Known peer behavior | The pinned Go and Rust implementations overlap with different parts of the profile, but the maintained-peer assessment finds no currently maintained independent implementation of the complete package contract. |
| Selected behavior | `verkletree-bandersnatch-ipa-256-v0` is the only implemented profile identity; its width, key and value layout, curve and quotient group, generator set, transcript, canonical encodings, tree topology, proof format, witness format, and state semantics are fixed and not caller-composable. |
| Rationale | A package-owned identity makes stored and exchanged bytes auditable without turning bounded corpus agreement into a general external compliance claim. |
| Security consequences | Unknown or inconsistent profile identities are rejected before cryptographic work, preventing cross-profile replay and attacker-selected component combinations. |
| Resource consequences | The selected profile retains the documented key, value, proof, witness, query, point-work, and temporary-memory limits; this decision introduces no new runtime input. |
| Compatibility consequences | The stable Go API is compatible only with the exact package-owned profile; changing any profile component requires a new decision, compatibility review, and versioned migration. |
| Wire consequences | Canonical package bytes carry the package profile identity and do not imply compatibility with generic Verkle, Ethereum, Go, Rust, database, or persistence formats. |
| Upstream status | No external standards body owns this package profile; the pinned Go implementation is in maintenance mode and the pinned Rust implementation has had no source update since 2024-10-25. |
| Reconsider when | A deliberately versioned replacement package profile or maintained independent implementation covers the complete contract with normative, hostile-input, and reproducible differential evidence. |

Credible interpretations:

- Claim compatibility with an external implementation from overlapping fixtures.
- Expose independently configurable cryptographic components.
- Freeze one package-owned research profile and limit interoperability claims to exact pinned evidence.

Executable evidence: `TestProfileIsExactAndRejectsOtherValues`,
`TestRustVerkleEncodingVectors`, `TestRustVerkleGeneratorSet`,
`TestCommitmentEngineMatchesPinnedRustVectors`,
`TestBuildMatchesPinnedRustTreeRoots`, and
`TestStatelessUpdaterMatchesPinnedRustRebuiltTransitions`.

Fixture evidence: `internal/backend/testdata/rust-verkle-encoding.tsv`,
`internal/backend/testdata/rust-verkle-generators.tsv`,
`internal/backend/testdata/rust-verkle-vector-commitments.tsv`,
`internal/committedtree/testdata/rust-verkle-tree-roots.tsv`, and
`internal/authstate/testdata/rust-verkle-transitions.tsv`.

Fuzz evidence: `FuzzDecodeTreeProof` and
`FuzzStatelessUpdaterMatchesStatefulTransition`. Interoperability and
differential evidence are intentionally empty because the retained corpora do
not qualify as official fixtures, provider evidence, or maintained-peer
differential evidence.

Public APIs: `Profile`, `BandersnatchIPA256V0`, and
roots, snapshots, proofs, witnesses, and their canonical encodings.

Documentation: `docs/specification-decisions.md`,
`specification/bandersnatch-ipa-256-v0.md`,
`specification/profile-freeze.md`, and
`specification/peer-assessment.json`.

## VERKLE-DEC-002: Ethereum Verkle EIPs are research inputs, not implemented profiles

| Field | Decision |
| --- | --- |
| Status and owner | `resolved`; `go-verkle-tree` maintainers |
| Classification and scope | `interoperability policy`; `application-policy` |
| Specification and version | `Ethereum Verkle EIPs 4762, 6800, 7612, and 7748 at c55786f4242e5324afd14c6bca890a369a771d7f (research only; not implemented)`; `EIPs c55786f4242e5324afd14c6bca890a369a771d7f (EIP-6800)` |
| Authority and section | `ethereum-eip-6800-source`; [EIP-6800](https://raw.githubusercontent.com/ethereum/EIPs/c55786f4242e5324afd14c6bca890a369a771d7f/EIPS/eip-6800.md); `EIP-6800 complete document; related EIPs 4762, 7612, and 7748`; `informative` |
| Issue | The cited Ethereum proposals describe protocol-specific state, gas, transition, and conversion work but do not constitute an activated stable profile implemented by this generic package. |
| Known peer behavior | At the pinned EIPs revision, EIPs 4762 and 7748 are Draft while EIPs 6800 and 7612 are Stagnant; current Geth direction replaces its Verkle work with a binary-tree design. |
| Selected behavior | EIPs 4762, 6800, 7612, and 7748 are unimplemented research inputs. The package does not implement their Ethereum state layout, gas accounting, overlay transition, historical conversion, activation, wire protocol, or mainnet compatibility. |
| Rationale | Separating research provenance from implemented behavior prevents cryptographic overlap and historical client work from being presented as protocol conformance. |
| Security consequences | Callers cannot rely on this package to enforce Ethereum-specific witness, gas, transition, or activation rules; doing so without a separate verified protocol layer would be unsafe. |
| Resource consequences | No Ethereum-specific gas schedule, state conversion, or protocol workload is accepted or budgeted by this package. |
| Compatibility consequences | Adopters requiring an Ethereum profile need a separately versioned implementation and migration evidence; the package-owned profile cannot be relabeled as an Ethereum profile. |
| Wire consequences | Package roots, proofs, witnesses, snapshots, and storage images are not Ethereum EIP wire objects and must not be exchanged as such. |
| Upstream status | The four proposals remain Draft or Stagnant at the pinned authority revision and no activated Ethereum Verkle profile is claimed by this package. |
| Reconsider when | Ethereum finalizes and activates a revision-pinned Verkle profile and a separately scoped implementation has complete normative, client, fixture, hostile-input, migration, and release evidence. |

Credible interpretations:

- Treat research lineage and overlapping tree primitives as Ethereum EIP compliance.
- Implement the EIPs as one implicit package profile.
- Record the exact proposals as monitored research authorities while explicitly excluding their protocol behavior from the implemented profile.

Additional authoritative sources:
`{"id":"ethereum-eip-4762-source","version":"EIPs c55786f4242e5324afd14c6bca890a369a771d7f (EIP-4762)","url":"https://raw.githubusercontent.com/ethereum/EIPs/c55786f4242e5324afd14c6bca890a369a771d7f/EIPS/eip-4762.md","specifications":["Ethereum Verkle EIPs 4762, 6800, 7612, and 7748 at c55786f4242e5324afd14c6bca890a369a771d7f (research only; not implemented)"]}`,
`{"id":"ethereum-eip-7612-source","version":"EIPs c55786f4242e5324afd14c6bca890a369a771d7f (EIP-7612)","url":"https://raw.githubusercontent.com/ethereum/EIPs/c55786f4242e5324afd14c6bca890a369a771d7f/EIPS/eip-7612.md","specifications":["Ethereum Verkle EIPs 4762, 6800, 7612, and 7748 at c55786f4242e5324afd14c6bca890a369a771d7f (research only; not implemented)"]}`,
and
`{"id":"ethereum-eip-7748-source","version":"EIPs c55786f4242e5324afd14c6bca890a369a771d7f (EIP-7748)","url":"https://raw.githubusercontent.com/ethereum/EIPs/c55786f4242e5324afd14c6bca890a369a771d7f/EIPS/eip-7748.md","specifications":["Ethereum Verkle EIPs 4762, 6800, 7612, and 7748 at c55786f4242e5324afd14c6bca890a369a771d7f (research only; not implemented)"]}`.

Executable evidence: `TestProfileIsExactAndRejectsOtherValues`. Fixture, fuzz,
interoperability, and differential evidence are empty because the decision is
an explicit exclusion rather than an Ethereum conformance implementation.

Public APIs: `Profile` and `BandersnatchIPA256V0`.

Documentation: `docs/specification-decisions.md`, `docs/compatibility.md`,
`docs/adoption.md`, and `specification/peer-assessment.json`.
