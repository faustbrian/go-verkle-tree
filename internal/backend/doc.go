// Package backend isolates the pinned Bandersnatch/Banderwagon IPA dependency
// behind profile-v0 encodings, explicit budgets and opaque values. Admission of
// canonical points and scalars is separate from proof verification. Identity
// proof elements are permitted; ordinary commitment encodings reject identity,
// and empty roots use an explicit container kind.
//
// This is an experimental research boundary, not a production cryptographic
// implementation. Setup and admitted proof calls cannot be interrupted inside
// the dependency. The operation gate is per engine, not process-wide; backend
// globals remain authoritative, and no constant-time guarantee is made.
package backend
