package emulator

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

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
