// Package verkletree provides an experimental, storage-independent authenticated
// key/value tree for the package-owned verkletree-bandersnatch-ipa-256-v0 profile.
// It is not production-ready and must not protect production integrity or
// confidentiality. Stable v1 API and canonical research encodings do not imply
// cryptographic soundness, a production audit, constant-time behavior, or
// Ethereum protocol compatibility; see SECURITY.md and docs/backend-audit.md.
//
// Keys and values are fixed 32-byte arrays. A present all-zero value differs
// from absence; Delete removes a key. Snapshots, roots, proofs, witnesses and
// operation results retain immutable state. Constructors copy caller data;
// callers must not mutate borrowed slices while a call is reading them.
// Canonical decoding and container construction admit representations, not
// trusted statements: use ProofEngine.VerifyForKeys and
// StatelessEngine.ApplyForRoot with application-trusted roots when authorizing
// reads or transitions. Lower-level verification accepts the embedded root.
//
// Expensive operations require explicit finite budgets and contexts. Scratch
// accounting is not a total process-memory cap. Proof operations serialize
// entry to the dependency per engine; queue waiting observes cancellation,
// but an admitted backend call cannot be interrupted. Independently constructed
// engines have independent gates. Pinned dependency globals remain authoritative
// and must not be mutated by callers; immutable facade values do not isolate
// cryptographic configuration from such mutation.
//
// Caller-owned storage adapters assert and implement atomicity, durability,
// compare-and-swap publication, isolated reads and maintenance. The package
// validates those declarations and verifies canonical nodes and reachable roots,
// but cannot enforce the adapter's persistence guarantees. Audit reports are
// not deletion authority; maintenance revalidates a fresh view and submits an
// atomic request. Recovery prunes unpublished node-only debris, not missing or
// corrupt published state. No concrete storage adapter is provided.
package verkletree
