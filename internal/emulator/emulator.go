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
	keypad     [16]byte
	screen     [64 * 32]bool
	rng        func() byte
}

func NewEmulator() Emulator {
	e := Emulator{
		pc:  START_ADDRESS,
		rng: NewSeededRng(1234),
		// rng: NewRng(),
	}

	copy(e.memory[50:], fontset[:])

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
