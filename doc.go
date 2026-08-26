// Package verkletree provides a stable v1 API for an explicitly profiled,
// storage-independent Verkle tree.
//
// The package exposes one package-owned research profile v0, immutable
// snapshots and roots, canonical whole-snapshot bytes, canonical atomic
// updates, and bounded aggregate membership and non-membership proofs. Every
// expensive operation requires a
// context and explicit resource limits. Snapshots can produce canonical
// content-addressed node batches for capability-checked atomic publication and
// reconstruct snapshots from capability-checked isolated reads after verifying
// every reachable node, root, and content address. A bounded audit can compare
// the complete canonical node inventory with all verified current and retained
// roots without mutating storage. A separate capability-checked operation can
// atomically replace the retained-publication set and prune only nodes outside
// the current and desired retained roots. Bounded recovery preserves every
// verified publication while atomically pruning node-only debris left by an
// interrupted unpublished write. Canonical stateless witnesses can
// verify authenticated Set operations on present, missing, or different stem
// paths and Delete operations that are absent, leave a stem non-empty, or
// remove stems and canonically collapse authenticated unary paths, then
// independently derive and match the claimed post-state root. Restoration of
// missing or corrupt published state and concrete storage adapters remain
// unavailable.
//
// The exported API is stable within v1 and exposes research profile v0 to
// evaluate a pinned commitment backend and complete tree semantics. Profile
// conformance does not imply production suitability, external audit, or
// Ethereum protocol compatibility.
package verkletree
