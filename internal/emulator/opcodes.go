package emulator

import "fmt"

var (
	ERROR_INVALID_OPCODE = fmt.Errorf("Invalid Opcode")
)

func (e *Emulator) decode(opcode word) (func(word), error) {

	var f func(word) = nil

	switch opcode & 0xf000 {
	case 0x0000:
		switch opcode & 0x0fff {
		case 0x00E0:
			f = e.op_00E0
		case 0x00EE:
			f = e.op_00EE
		}
	case 0x1000:
		f = e.op_1nnn
	case 0x2000:
		f = e.op_2nnn
	case 0x3000:
		f = e.op_3xkk
	case 0x4000:
		f = e.op_4xkk
	case 0x5000:
		f = e.op_5xy0
	case 0x6000:
		f = e.op_6xkk
	case 0x7000:
		f = e.op_7xkk
	case 0x8000:
		switch opcode & 0xf {
		case 0x0:
			f = e.op_8xy0
		case 0x1:
			f = e.op_8xy1
		case 0x2:
			f = e.op_8xy2
		case 0x3:
			f = e.op_8xy3
		case 0x4:
			f = e.op_8xy4
		case 0x5:
			f = e.op_8xy5
		case 0x6:
			f = e.op_8xy6
		case 0x7:
			f = e.op_8xy7
		case 0xE:
			f = e.op_8xyE
		}
	case 0x9000:
		f = e.op_9xy0
	case 0xA000:
		f = e.op_Annn
	case 0xB000:
		f = e.op_Bnnn
	case 0xC000:
		f = e.op_Cxkk
	case 0xD000:
		f = e.op_Dxyn
	case 0xE000:
		switch opcode & 0xff {
		case 0xa1:
			f = e.op_ExA1
		case 0x9e:
			f = e.op_Ex9E
		}
	case 0xF000:
		switch opcode & 0xff {

		case 0x07:
			f = e.op_Fx07
		case 0x0A:
			f = e.op_Fx0A
		case 0x15:
			f = e.op_Fx15
		case 0x18:
			f = e.op_Fx18
		case 0x1E:
			f = e.op_Fx1E
		case 0x29:
			f = e.op_Fx29
		case 0x33:
			f = e.op_Fx33
		case 0x55:
			f = e.op_Fx55
		}
	}

	if f == nil {
		return f, fmt.Errorf("%w : %v", ERROR_INVALID_OPCODE, opcode)
	}

	return f, nil
}

// Clear the display
func (e *Emulator) op_00E0(opcode word) {
	for i := range len(e.screen) {
		e.screen[i] = false
	}
}

func (e *Emulator) op_00EE(opcode word) {

}

func (e *Emulator) op_1nnn(opcode word) {

}

func (e *Emulator) op_2nnn(opcode word) {

}

func (e *Emulator) op_3xkk(opcode word) {

}

func (e *Emulator) op_4xkk(opcode word) {

}

func (e *Emulator) op_5xy0(opcode word) {

}

func (e *Emulator) op_6xkk(opcode word) {

}

func (e *Emulator) op_7xkk(opcode word) {

}

func (e *Emulator) op_8xy0(opcode word) {

}

func (e *Emulator) op_8xy1(opcode word) {

}

func (e *Emulator) op_8xy2(opcode word) {

}

func (e *Emulator) op_8xy3(opcode word) {

}

func (e *Emulator) op_8xy4(opcode word) {

}

func (e *Emulator) op_8xy5(opcode word) {

}

func (e *Emulator) op_8xy6(opcode word) {

}

func (e *Emulator) op_8xy7(opcode word) {

}

func (e *Emulator) op_8xyE(opcode word) {

}

func (e *Emulator) op_9xy0(opcode word) {

}

func (e *Emulator) op_Annn(opcode word) {

}

func (e *Emulator) op_Bnnn(opcode word) {

}

func (e *Emulator) op_Cxkk(opcode word) {

}

func (e *Emulator) op_Dxyn(opcode word) {

}

func (e *Emulator) op_Ex9E(opcode word) {

}

func (e *Emulator) op_ExA1(opcode word) {

}

func (e *Emulator) op_Fx07(opcode word) {

}

func (e *Emulator) op_Fx0A(opcode word) {

}

func (e *Emulator) op_Fx15(opcode word) {

}

func (e *Emulator) op_Fx18(opcode word) {

}

func (e *Emulator) op_Fx1E(opcode word) {

}

func (e *Emulator) op_Fx29(opcode word) {

}

func (e *Emulator) op_Fx33(opcode word) {

}

func (e *Emulator) op_Fx55(opcode word) {

}

func (e *Emulator) op_Fx65(opcode word) {

}
