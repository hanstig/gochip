package emulator

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSomething(t *testing.T) {
	assert.Equal(t, 123, 123, "they should be equal")
	require.NotEqual(t, 123, 456, "they should not be equal")
}
