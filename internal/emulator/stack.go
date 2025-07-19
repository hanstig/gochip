package emulator

import (
	"fmt"
)

const STACK_SIZE = 16

var (
	ERROR_IS_EMPTY = fmt.Errorf("The stack is empty!")
	ERROR_IS_FULL  = fmt.Errorf("The stack is full!")
)

type Stack struct {
	data [STACK_SIZE]uint16
	size uint8
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) IsEmpty() bool {
	return s.size == 0
}

func (s *Stack) IsFull() bool {
	return s.size == STACK_SIZE
}

func (s *Stack) Size() uint8 {
	return s.size
}

func (s *Stack) Push(v uint16) error {
	if s.IsFull() {
		return ERROR_IS_FULL
	}
	s.data[s.size] = v
	s.size++
	return nil
}

func (s *Stack) Pop() (uint16, error) {
	if s.IsEmpty() {
		return 0, ERROR_IS_EMPTY
	}
	s.size--
	return s.data[s.size], nil
}
