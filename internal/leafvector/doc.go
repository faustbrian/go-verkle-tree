// Package leafvector defines research profile-v0 field inputs for fixed stems
// and values without constructing commitments. A suffix selects two positions
// in one of two 256-wide vectors. Present low halves carry a 2^128 marker so a
// present all-zero value differs from the two zero scalars denoting absence.
// These bounded little-endian inputs need no scalar-field modular reduction.
package leafvector
