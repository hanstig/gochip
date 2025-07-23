package emulator

import (
	"fmt"
	"io"
	"os"
)

type word = uint16

const (
	START_ADDRESS         word = 0x200
	FONTSET_START_ADDRESS word = 0x50
	SCREEN_WIDTH          int  = 64
	SCREEN_HEIGHT         int  = 32
)

var (
	ERROR_LOADING_ROM = fmt.Errorf("Error loading rom")
)

type Emulator struct {
	registers  [16]byte
	memory     [4096]byte
	index      word
	pc         word
	stack      [16]word
	sp         byte
	delayTimer byte
	soundTimer byte
	Keypad     [16]bool
	Screen     [SCREEN_HEIGHT * SCREEN_WIDTH]bool
	rng        func() byte
}

func NewSeededEmulator(seed uint64) Emulator {
	e := NewEmulator()
	e.rng = NewSeededRng(seed)
	return e
}

func NewEmulator() Emulator {
	e := Emulator{
		pc:  START_ADDRESS,
		rng: NewRng(),
	}
	copy(e.memory[FONTSET_START_ADDRESS:], fontset[:])
	return e
}

func (e *Emulator) LoadROM(filepath string) error {

	f, err := os.Open(filepath)
	if err != nil {
		return fmt.Errorf("%v: %w", ERROR_LOADING_ROM, err)
	}
	defer f.Close()

	_, err = f.Read(e.memory[START_ADDRESS:])

	if err != nil {
		return fmt.Errorf("%v: %w", ERROR_LOADING_ROM, err)
	}

	extra := make([]byte, 1)
	_, err = f.Read(extra)
	if err != io.EOF {
		return fmt.Errorf("%v: ROM file couldnt fit into memory!", ERROR_LOADING_ROM)

	}

	return nil
}

func (e *Emulator) Cycle() error {
	opcode := (word(e.memory[e.pc]) << 8) | word(e.memory[e.pc+1])
	e.pc += 2

	f, err := e.decode(opcode)

	if err != nil {
		return err
	}

	f(opcode)

	if e.delayTimer > 0 {
		e.delayTimer--
	}

	if e.soundTimer > 0 {
		e.soundTimer--
	}

	return nil
}
