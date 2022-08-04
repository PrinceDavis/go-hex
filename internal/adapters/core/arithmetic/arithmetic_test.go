package arithmetic

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAddition(t *testing.T) {
	arith := NewAdapter()

	answer, err := arith.Addition(1, 2)
	if err != nil {
		t.Fatalf("expected: %v,  got: %v", nil, err)
	}
	require.Equal(t, int32(3), answer)
}
