package emulator

import "math/rand/v2"

func NewSeededRng(seed uint64) func() byte {
	r := rand.New(rand.NewPCG(seed, seed+1))
	return func() byte {
		return byte(r.Int32N(256))
	}
}

func NewRng() func() byte {
	return func() byte {
		return byte(rand.Int32N(256))
	}
}
