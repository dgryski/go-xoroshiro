// Package xoroshiro implements the xoroshiro128+ RNG
/*

This is a literal translation of http://xoroshiro.di.unimi.it/xoroshiro128plus.c
and http://xoroshiro.di.unimi.it/splitmix64.c

*/
package xoroshiro

// State is a xoroshiro128+ RNG state.
type State [2]uint64

// New returns a new RNG
func New(seed int64) State {
	var s State
	s.Seed(seed)
	return s
}

// Next returns the next number in the sequence
func (s *State) Next() uint64 {
	s0, s1 := s[0], s[1]
	result := s0 + s1

	s1 ^= s0
	s[0] = rotl(s0, 55) ^ s1 ^ (s1 << 14) // a, b
	s[1] = rotl(s1, 36)                   // c

	return result
}

func (s *State) Int63() int64 {
	return int64(s.Next() & 0x7fffffffffffffff)
}

// Int63n returns a uniform integer [0, n)
func (s *State) Int63n(n int64) int64 {
	if n <= 0 {
		panic("invalid argument to Int63n")
	}
	// shortcut powers of 2
	if n&(n-1) == 0 {
		return s.Int63() & (n - 1)
	}
	max := int64((1 << 63) - 1 - (1<<63)%uint64(n))
	r := s.Int63()
	for r > max {
		r = s.Int63()
	}
	return r % n
}

func (s *State) Seed(seed int64) {
	splitmix := SplitMix64(seed)
	s[0], s[1] = splitmix.Next(), splitmix.Next()
}

func (s *State) Jump() {
	var JUMP = [2]uint64{0xbeac0467eba5facb, 0xd86b048b86aa9922}

	var s0, s1 uint64

	for _, v := range JUMP {
		for b := uint(0); b < 64; b++ {
			if (v & (1 << b)) != 0 {
				s0 ^= s[0]
				s1 ^= s[1]
			}
			s.Next()
		}
	}

	s[0] = s0
	s[1] = s1
}

func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

// SplitMix64 is an RNG with 64-bits of state
type SplitMix64 uint64

// Next returns the next number in the sequence
func (x *SplitMix64) Next() uint64 {
	*x += 0x9E3779B97F4A7C15
	z := uint64(*x)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	z = z ^ (z >> 31)
	return z
}
