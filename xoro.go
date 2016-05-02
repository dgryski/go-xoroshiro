package xoro

type State [2]uint64

func (s *State) Next() uint64 {
	s0, s1 := s[0], s[1]
	result := s0 + s1

	s1 ^= s0
	s[0] = rotl(s0, 55) ^ s1 ^ (s1 << 14) // a, b
	s[1] = rotl(s1, 36)                   // c

	return result
}

func rotl(x uint64, k uint) uint64 {
	return (x << k) | (x >> (64 - k))
}

type ShiftMix64 uint64

func (x *ShiftMix64) Next() uint64 {
	*x += 0x9E3779B97F4A7C15
	z := uint64(*x)
	z = (z ^ (z >> 30)) * 0xBF58476D1CE4E5B9
	z = (z ^ (z >> 27)) * 0x94D049BB133111EB
	z = z ^ (z >> 31)
	return z
}
