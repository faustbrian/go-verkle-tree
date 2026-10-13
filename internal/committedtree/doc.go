// Package committedtree constructs immutable research profile-v0 node arenas,
// vector commitments, proof material and content-addressed storage encodings.
// Construction copies entries and orders raw keys; array values distinguish
// present zero from absence. Storage decoding admits local node structure;
// complete topology and commitment reconstruction are verified by the loader.
package committedtree
