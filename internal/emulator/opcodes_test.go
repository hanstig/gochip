package emulator

import (
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func FuzzDecoder(f *testing.F) {
// 	// Valid op_1nnn
// 	testcases := []word{0x000, 0xfff, 0x1}
// 	for _, tc := range testcases {
// 		f.Add(tc)
// 	}
//
// 	f.Fuzz(func(t *testing.T, orig word) {
// 		em := NewEmulator()
// 		opcode := 0x1000 | (orig & 0x0fff)
//
// 		f, err := em.decode(opcode)
//
// 		assert.Nil(t, err, "Gave error for valid opcode")
// 		expected := reflect.ValueOf(em.op_1nnn).Pointer()
// 		actual := reflect.ValueOf(f).Pointer()
//
// 		assert.Equal(t, expected, actual, "Should return pointer to op_1nnn")
// 	})
// }

func TestDecoder(t *testing.T) {
	em := NewEmulator()

	f, err := em.decode(0x1000)
	assert.Nil(t, err, "no err valid opcode")

	expected := reflect.ValueOf(em.op_1nnn).Pointer()
	actual := reflect.ValueOf(f).Pointer()
	assert.Equal(t, expected, actual, "Returned correct function pointer")

	f, err = em.decode(0x0e00)
	assert.NotNil(t, err, "decoding invalid opcode")
}

func TestOpcode_00E0(t *testing.T) {
	empty := [64 * 32]bool{}

	em := NewEmulator()
	em.op_00E0(0x00E0)

	assert.True(t, slices.Equal(em.screen[:], empty[:]), "fresh screen empty after cls")

	for i := range len(em.screen) {
		em.screen[i] = true
	}

	assert.False(t, slices.Equal(em.screen[:], empty[:]), "Sanity test, screen not empty")
	em.op_00E0(0x00E0)
	assert.True(t, slices.Equal(em.screen[:], empty[:]), "Full screen becomes empty after cls")

}
