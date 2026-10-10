package verkletree

import "testing"

// These arithmetic-only cases allocate no node storage. The supported count
// domain is MaxInt32 even on wider hosts; CI also executes them on linux/386.
func TestStorageAuditResultCapacityPreserves32BitBounds(t *testing.T) {
	t.Parallel()
	const maximum = 1<<31 - 1
	for _, test := range []struct {
		name                         string
		length, capacity, additional int
	}{
		{"rounded power of two", 0, 0, 1<<30 + 1},
		{"required count addition", maximum - 1, maximum - 1, 2},
		{"rounding numerator addition", 1 << 30, 1 << 30, 1},
	} {
		t.Run(test.name, func(t *testing.T) {
			got := storageAuditResultCapacity(test.length, test.capacity, test.additional, maximum)
			if got != maximum {
				t.Fatalf("result capacity = %d, want %d", got, maximum)
			}
		})
	}
}
