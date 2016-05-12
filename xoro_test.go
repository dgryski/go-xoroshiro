package xoroshiro

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestSplitMix(t *testing.T) {

	var tests = []struct {
		seed  uint64
		nexts []uint64
	}{
		{0xdeadbeef, []uint64{0x4adfb90f68c9eb9b, 0xde586a3141a10922, 0x21fbc2f8e1cfc1d, 0x7466ce737be16790, 0x3bfa8764f685bd1c, 0xab203e503cb55b3f, 0x5a2fdc2bf68cedb3, 0xb30a4ccf430b1b5a, 0xa90415039bd5985, 0x26ae50847745eb7e}},
	}

	for _, tt := range tests {
		s := SplitMix64(tt.seed)
		for i, n := range tt.nexts {
			if got := s.Next(); got != n {
				t.Errorf("SplitMix(%d, %d).Next()=%x, want %x\n", tt.seed, i, got, n)
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

func TestXoroJump(t *testing.T) {

	var tests = []struct {
		seed  [2]uint64
		nexts []uint64
	}{
		{[2]uint64{0x916df851e2aee44, 0x9ade0f09ffca1bc4},
			[]uint64{
				7317131579098254132,
				9124900356304480981,
				16687222659825326268,
				10655786156111842618,
				384402176967881600,
				2173327412138143738,
				14504858356897473757,
				4786136656534720403,
				3081009357741310655,
				4072612981517571462,
			},
		},
	}

	for _, tt := range tests {
		s := State(tt.seed)
		for i, n := range tt.nexts {
			s.Jump()
			if got := s.Next(); got != n {
				t.Errorf("XoroShiro(%v, %d).Jump().Next()=%x, want %x\n", tt.seed, i, got, n)
			}
		}
	}
}

func ExampleState() {

	// If you have 128 bits of randomness for a seed, use those.
	// If you have a single 64-bit seed, generate more with SplitMix64.

	seed := SplitMix64(0x0ddc0ffeebadf00d)

	s := State([2]uint64{seed.Next(), seed.Next()})

	for i := 0; i < 10; i++ {
		fmt.Println(s.Next())
	}

	// Output:
	// 11814330020949985800
	// 11817088786836023749
	// 1654166990350674155
	// 14112748191344281834
	// 4288295283113472773
	// 8391955421631067594
	// 168274855724945977
	// 2815117763357611551
	// 12187186948608395331
	// 10629044371437376348
}

// Verify State implements rand.Source
var _ rand.Source = &State{}
