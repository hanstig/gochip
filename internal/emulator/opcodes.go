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
		case 0x65:
			f = e.op_Fx65
		}
	}

	if f == nil {
		return f, fmt.Errorf("%w : %v", ERROR_INVALID_OPCODE, opcode)
	}

	return f, nil
}

// Clear the display
func (e *Emulator) op_00E0(opcode word) {
	for i := range len(e.Screen) {
		e.Screen[i] = false
	}
}

// Return from a subroutine
func (e *Emulator) op_00EE(opcode word) {
	if e.sp == 0 {
		return
	}
	e.sp--
	e.pc = e.stack[e.sp]
}

// Jump to nnn
func (e *Emulator) op_1nnn(opcode word) {
	e.pc = opcode & 0x0fff
}

// Call subroutine at nnn
func (e *Emulator) op_2nnn(opcode word) {
	e.stack[e.sp] = e.pc
	e.sp++
	e.pc = opcode & 0xfff
}

// Skip next instruction if Vx == kk
func (e *Emulator) op_3xkk(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	kk := byte(opcode & 0xff)

	if e.registers[Vx] == kk {
		e.pc += 2
	}
}

// Skip next instruction if Vx != kk
func (e *Emulator) op_4xkk(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	kk := byte(opcode & 0xff)

	if e.registers[Vx] != kk {
		e.pc += 2
	}
}

// Skip next instruction if Vx == Vy
func (e *Emulator) op_5xy0(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)

	if e.registers[Vx] == e.registers[Vy] {
		e.pc += 2
	}
}

// Set Vx = kk
func (e *Emulator) op_6xkk(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	kk := byte(opcode & 0xff)
	e.registers[Vx] = kk
}

// Set Vx = Vx + kk
func (e *Emulator) op_7xkk(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	kk := byte(opcode & 0xff)

	//NOTE: Correct overflow behaviour?
	e.registers[Vx] += kk
}

// Set Vx = Vy
func (e *Emulator) op_8xy0(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)
	e.registers[Vx] = e.registers[Vy]
}

// Set Vx = Vx | Vy
func (e *Emulator) op_8xy1(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)
	e.registers[Vx] |= e.registers[Vy]
}

// Set Vx = Vx & Vy
func (e *Emulator) op_8xy2(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)
	e.registers[Vx] &= e.registers[Vy]
}

// Set Vx = Vx ^ Vy
func (e *Emulator) op_8xy3(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)
	e.registers[Vx] ^= e.registers[Vy]
}

// Set Vx = Vx 0 Vy, Set VF = carry
func (e *Emulator) op_8xy4(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)

	sum := word(e.registers[Vx]) + word(e.registers[Vy])
	if sum > 255 {
		e.registers[0xf] = 1
	} else {
		e.registers[0xf] = 0
	}

	e.registers[Vx] = byte(sum & 0xff)
}

// Set Vx = Vx - Vy, set VF = not underflow
func (e *Emulator) op_8xy5(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)

	if e.registers[Vx] > e.registers[Vy] {
		e.registers[0xf] = 1
	} else {
		e.registers[0xf] = 0
	}

	e.registers[Vx] -= e.registers[Vy]
}

// Save least significant bit in Vf, then shif vx right 1
func (e *Emulator) op_8xy6(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	e.registers[0xf] = e.registers[Vx] & 0x1
	e.registers[Vx] >>= 1
}

// Set Vx = Vy - Vx, Vf = not underflow
func (e *Emulator) op_8xy7(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)
	if e.registers[Vx] < e.registers[Vy] {
		e.registers[0xf] = 1
	} else {
		e.registers[0xf] = 0
	}
	e.registers[Vx] = e.registers[Vy] - e.registers[Vx]
}

// Save most significant bit in Vf, then shif vx left 1
func (e *Emulator) op_8xyE(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)

	e.registers[0xf] = (e.registers[Vx] & 0x80) >> 7
	e.registers[Vx] <<= 1
}

// Skip next instruction if Vx != Vy
func (e *Emulator) op_9xy0(opcode word) {
	Vx := byte((opcode & 0x0f00) >> 8)
	Vy := byte((opcode & 0x00f0) >> 4)
	if e.registers[Vx] != e.registers[Vy] {
		e.pc += 2
	}
}

// set index = nnn
func (e *Emulator) op_Annn(opcode word) {
	e.index = opcode & 0x0fff
}

// Jump to V0 + nnn
func (e *Emulator) op_Bnnn(opcode word) {
	e.pc = word(e.registers[0]) + (opcode & 0x0fff)
}

// set Vx = random & kk
func (e *Emulator) op_Cxkk(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	kk := byte(opcode & 0x00ff)

	e.registers[Vx] = e.rng() & kk
}

// draw sprite
func (e *Emulator) op_Dxyn(opcode word) {
	x := int(e.registers[(opcode&0x0F00)>>8])
	y := int(e.registers[(opcode&0x00F0)>>4])
	h := opcode & 0x000F

	e.registers[0xF] = 0
	for row := word(0) ; row < h && y + int(row) < SCREEN_HEIGHT; row++{
		spriteRow := e.memory[e.index+row]
		for col := 0 ; col < 8 && x + col < SCREEN_WIDTH; col++ {
			spritePixel := (spriteRow & (0x80 >> col)) != 0
			screenPixel := &e.Screen[(y+int(row))*SCREEN_WIDTH+x+int(col)]

			if spritePixel {
				if *screenPixel {
					e.registers[0xf] = 1
				}
				*screenPixel = !*screenPixel
			}

		}
	}
}

// Skip instruction if key[Vx] is pressed
func (e *Emulator) op_Ex9E(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	keyNum := e.registers[Vx]
	if e.Keypad[keyNum] {
		e.pc += 2
	}
}

// Skip instruction if key[Vx] is not pressed
func (e *Emulator) op_ExA1(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	keyNum := e.registers[Vx]
	if !e.Keypad[keyNum] {
		e.pc += 2
	}
}

// Set Vx to delay timer value
func (e *Emulator) op_Fx07(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	e.registers[Vx] = e.delayTimer

}

// Wait for keypress, then store value in vx
func (e *Emulator) op_Fx0A(opcode word) {
	Vx := (opcode & 0x0f00) >> 8

	for i, v := range e.Keypad {
		if v {
			e.registers[Vx] = byte(i)
			return
		}
	}
	e.pc -= 2
}

// Set delaytimer to Vx
func (e *Emulator) op_Fx15(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	e.delayTimer = e.registers[Vx]
}

// Set sound timer to Vx
func (e *Emulator) op_Fx18(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	e.soundTimer = e.registers[Vx]
}

// Set index += Vx
func (e *Emulator) op_Fx1E(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	e.index += word(e.registers[Vx])
}

// Set I to sprite for digit in Vx
func (e *Emulator) op_Fx29(opcode word) {
	Vx := (opcode & 0x0f00) >> 8
	digit := e.registers[Vx]
	e.index = FONTSET_START_ADDRESS + (5 * word(digit))
}

// save Bcd of Vx at I, I+1, I+2
func (e *Emulator) op_Fx33(opcode word) {
	Vx := (opcode & 0x0F00) >> 8
	value := e.registers[Vx]

	e.memory[e.index+2] = value % 10
	value /= 10

	e.memory[e.index+1] = value % 10
	value /= 10

	e.memory[e.index] = value % 10
}

// Store registers V0-Vx in memory from index location forward
func (e *Emulator) op_Fx55(opcode word) {
	Vx := (opcode & 0x0F00) >> 8
	for i := word(0); i <= Vx; i++ {
		e.memory[e.index+i] = e.registers[i]
	}
}

// Read registers V0-Vx from memory at index location forward
func (e *Emulator) op_Fx65(opcode word) {
	Vx := (opcode & 0x0F00) >> 8

	for i := word(0); i <= Vx; i++ {
		e.registers[i] = e.memory[e.index+i]
	}
}
