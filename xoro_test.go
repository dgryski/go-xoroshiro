package xoroshiro

import (
	"testing"
)

func TestShiftMix(t *testing.T) {

	var tests = []struct {
		seed  uint64
		nexts []uint64
	}{
		{0xdeadbeef, []uint64{0x4adfb90f68c9eb9b, 0xde586a3141a10922, 0x21fbc2f8e1cfc1d, 0x7466ce737be16790, 0x3bfa8764f685bd1c, 0xab203e503cb55b3f, 0x5a2fdc2bf68cedb3, 0xb30a4ccf430b1b5a, 0xa90415039bd5985, 0x26ae50847745eb7e}},
	}

	for _, tt := range tests {
		s := ShiftMix64(tt.seed)
		for i, n := range tt.nexts {
			if got := s.Next(); got != n {
				t.Errorf("ShiftMix(%d, %d).Next()=%x, want %x\n", tt.seed, i, got, n)
			}
		}
	}
}

func TestXoroShift(t *testing.T) {

	var tests = []struct {
		seed  [2]uint64
		nexts []uint64
	}{
		{[2]uint64{1, 2}, []uint64{0x3, 0x8000300000c003, 0x118406038000363, 0xa080fe5030c4c366, 0x3ae0e84f181c8404, 0x390283917940944, 0x98dcc1f06360888c, 0x7db94a025d95c80f, 0x775088046d70b290, 0x412422d94084790d}},
	}

	for _, tt := range tests {
		s := State(tt.seed)
		for i, n := range tt.nexts {
			if got := s.Next(); got != n {
				t.Errorf("XoroShiro(%v, %d).Next()=%x, want %x\n", tt.seed, i, got, n)
			}
		}
	}
}
