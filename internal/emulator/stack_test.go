package emulator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	st := NewStack()

	assert.Equal(t, uint8(0), st.Size(), "Should be empty on init")

	_, err := st.Pop()

	assert.NotNil(t, err, "Can't Pop empty stack")

	err = st.Push(1)
	assert.Nil(t, err, "Can Push to stack")
	st.Push(2)
	st.Push(3)
	st.Push(4)

	assert.Equal(t, st.Size(), uint8(4), "Size updates after push")

	val, err := st.Pop()
	assert.Nil(t, err, "Can pop from non-empty stack")
	assert.Equal(t, val, uint16(4), "Get correct val from popping stack")
	assert.Equal(t, st.Size(), uint8(3), "Size updates after pop")

	st = NewStack()

	for i := range STACK_SIZE {
		err = st.Push(uint16(i))
	}

	assert.Nil(t, err, "Can fill stack completely")

	st.Push(0)
	st.Push(0)
	st.Push(0)
	err = st.Push(0)

	assert.NotNil(t, err, "Cant overfill stack")
	assert.Equal(t, st.Size(), uint8(STACK_SIZE), "Pushing to full doesnt change size")

	val, err = st.Pop()
	assert.Equal(t, uint16(STACK_SIZE-1), val, "Pushing to full leaves top value alone")
}
