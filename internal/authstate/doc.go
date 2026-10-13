// Package authstate binds immutable ordered key/value state to the research
// profile-v0 committed tree, canonical claims, proofs and stateless witnesses.
// Container validity is not cryptographic verification or application root
// authority. Stateless updates authenticate disclosed old values before
// applying commitment deltas; topology-changing deletes require completeness
// probes to derive canonical collapse without undisclosed state.
package authstate
