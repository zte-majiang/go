// Copyright 2017 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bits

import (
	"testing"
	"unsafe"
)

func TestUintSize(t *testing.T) {
	var x uint
	if want := unsafe.Sizeof(x) * 8; UintSize != want {
		t.Fatalf("UintSize = %d; want %d", UintSize, want)
	}
}

func TestLeadingZeros(t *testing.T) {
	for i := 0; i < 256; i++ {
		nlz := tab[i].nlz
		for k := 0; k < 64-8; k++ {
			x := uint64(i) << uint(k)
			if x <= 1<<8-1 {
				got := LeadingZeros8(uint8(x))
				want := nlz - k + (8 - 8)
				if x == 0 {
					want = 8
				}
				if got != want {
					t.Fatalf("LeadingZeros8(%#02x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<16-1 {
				got := LeadingZeros16(uint16(x))
				want := nlz - k + (16 - 8)
				if x == 0 {
					want = 16
				}
				if got != want {
					t.Fatalf("LeadingZeros16(%#04x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<32-1 {
				got := LeadingZeros32(uint32(x))
				want := nlz - k + (32 - 8)
				if x == 0 {
					want = 32
				}
				if got != want {
					t.Fatalf("LeadingZeros32(%#08x) == %d; want %d", x, got, want)
				}
				if UintSize == 32 {
					got = LeadingZeros(uint(x))
					if got != want {
						t.Fatalf("LeadingZeros(%#08x) == %d; want %d", x, got, want)
					}
				}
			}

			if x <= 1<<64-1 {
				got := LeadingZeros64(uint64(x))
				want := nlz - k + (64 - 8)
				if x == 0 {
					want = 64
				}
				if got != want {
					t.Fatalf("LeadingZeros64(%#016x) == %d; want %d", x, got, want)
				}
				if UintSize == 64 {
					got = LeadingZeros(uint(x))
					if got != want {
						t.Fatalf("LeadingZeros(%#016x) == %d; want %d", x, got, want)
					}
				}
			}
		}
	}
}

func TestTrailingZeros(t *testing.T) {
	for i := 0; i < 256; i++ {
		ntz := tab[i].ntz
		for k := 0; k < 64-8; k++ {
			x := uint64(i) << uint(k)
			want := ntz + k
			if x <= 1<<8-1 {
				got := TrailingZeros8(uint8(x))
				if x == 0 {
					want = 8
				}
				if got != want {
					t.Fatalf("TrailingZeros8(%#02x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<16-1 {
				got := TrailingZeros16(uint16(x))
				if x == 0 {
					want = 16
				}
				if got != want {
					t.Fatalf("TrailingZeros16(%#04x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<32-1 {
				got := TrailingZeros32(uint32(x))
				if x == 0 {
					want = 32
				}
				if got != want {
					t.Fatalf("TrailingZeros32(%#08x) == %d; want %d", x, got, want)
				}
				if UintSize == 32 {
					got = TrailingZeros(uint(x))
					if got != want {
						t.Fatalf("TrailingZeros(%#08x) == %d; want %d", x, got, want)
					}
				}
			}

			if x <= 1<<64-1 {
				got := TrailingZeros64(uint64(x))
				if x == 0 {
					want = 64
				}
				if got != want {
					t.Fatalf("TrailingZeros64(%#016x) == %d; want %d", x, got, want)
				}
				if UintSize == 64 {
					got = TrailingZeros(uint(x))
					if got != want {
						t.Fatalf("TrailingZeros(%#016x) == %d; want %d", x, got, want)
					}
				}
			}
		}
	}
}

func TestOnesCount(t *testing.T) {
	for i := 0; i < 256; i++ {
		want := tab[i].pop
		for k := 0; k < 64-8; k++ {
			x := uint64(i) << uint(k)
			if x <= 1<<8-1 {
				got := OnesCount8(uint8(x))
				if got != want {
					t.Fatalf("OnesCount8(%#02x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<16-1 {
				got := OnesCount16(uint16(x))
				if got != want {
					t.Fatalf("OnesCount16(%#04x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<32-1 {
				got := OnesCount32(uint32(x))
				if got != want {
					t.Fatalf("OnesCount32(%#08x) == %d; want %d", x, got, want)
				}
				if UintSize == 32 {
					got = OnesCount(uint(x))
					if got != want {
						t.Fatalf("OnesCount(%#08x) == %d; want %d", x, got, want)
					}
				}
			}

			if x <= 1<<64-1 {
				got := OnesCount64(uint64(x))
				if got != want {
					t.Fatalf("OnesCount64(%#016x) == %d; want %d", x, got, want)
				}
				if UintSize == 64 {
					got = OnesCount(uint(x))
					if got != want {
						t.Fatalf("OnesCount(%#016x) == %d; want %d", x, got, want)
					}
				}
			}
		}
	}
}

func TestRotateLeft(t *testing.T) {
	var m uint64 = deBruijn64

	for k := uint(0); k < 128; k++ {
		x8 := uint8(m)
		got8 := RotateLeft8(x8, int(k))
		want8 := x8<<(k&0x7) | x8>>(8-k&0x7)
		if got8 != want8 {
			t.Fatalf("RotateLeft8(%#02x, %d) == %#02x; want %#02x", x8, k, got8, want8)
		}

		x16 := uint16(m)
		got16 := RotateLeft16(x16, int(k))
		want16 := x16<<(k&0xf) | x16>>(16-k&0xf)
		if got16 != want16 {
			t.Fatalf("RotateLeft16(%#04x, %d) == %#04x; want %#04x", x16, k, got16, want16)
		}

		x32 := uint32(m)
		got32 := RotateLeft32(x32, int(k))
		want32 := x32<<(k&0x1f) | x32>>(32-k&0x1f)
		if got32 != want32 {
			t.Fatalf("RotateLeft32(%#08x, %d) == %#08x; want %#08x", x32, k, got32, want32)
		}
		if UintSize == 32 {
			x := uint(m)
			got := RotateLeft(x, int(k))
			want := x<<(k&0x1f) | x>>(32-k&0x1f)
			if got != want {
				t.Fatalf("RotateLeft(%#08x, %d) == %#08x; want %#08x", x, k, got, want)
			}
		}

		x64 := uint64(m)
		got64 := RotateLeft64(x64, int(k))
		want64 := x64<<(k&0x3f) | x64>>(64-k&0x3f)
		if got64 != want64 {
			t.Fatalf("RotateLeft64(%#016x, %d) == %#016x; want %#016x", x64, k, got64, want64)
		}
		if UintSize == 64 {
			x := uint(m)
			got := RotateLeft(x, int(k))
			want := x<<(k&0x3f) | x>>(64-k&0x3f)
			if got != want {
				t.Fatalf("RotateLeft(%#016x, %d) == %#016x; want %#016x", x, k, got, want)
			}
		}
	}
}

func TestRotateRight(t *testing.T) {
	var m uint64 = deBruijn64

	for k := uint(0); k < 128; k++ {
		x8 := uint8(m)
		got8 := RotateRight8(x8, int(k))
		want8 := x8>>(k&0x7) | x8<<(8-k&0x7)
		if got8 != want8 {
			t.Fatalf("RotateRight8(%#02x, %d) == %#02x; want %#02x", x8, k, got8, want8)
		}

		x16 := uint16(m)
		got16 := RotateRight16(x16, int(k))
		want16 := x16>>(k&0xf) | x16<<(16-k&0xf)
		if got16 != want16 {
			t.Fatalf("RotateRight16(%#04x, %d) == %#04x; want %#04x", x16, k, got16, want16)
		}

		x32 := uint32(m)
		got32 := RotateRight32(x32, int(k))
		want32 := x32>>(k&0x1f) | x32<<(32-k&0x1f)
		if got32 != want32 {
			t.Fatalf("RotateRight32(%#08x, %d) == %#08x; want %#08x", x32, k, got32, want32)
		}
		if UintSize == 32 {
			x := uint(m)
			got := RotateRight(x, int(k))
			want := x>>(k&0x1f) | x<<(32-k&0x1f)
			if got != want {
				t.Fatalf("RotateRight(%#08x, %d) == %#08x; want %#08x", x, k, got, want)
			}
		}

		x64 := uint64(m)
		got64 := RotateRight64(x64, int(k))
		want64 := x64>>(k&0x3f) | x64<<(64-k&0x3f)
		if got64 != want64 {
			t.Fatalf("RotateRight64(%#016x, %d) == %#016x; want %#016x", x64, k, got64, want64)
		}
		if UintSize == 64 {
			x := uint(m)
			got := RotateRight(x, int(k))
			want := x>>(k&0x3f) | x<<(64-k&0x3f)
			if got != want {
				t.Fatalf("RotateRight(%#016x, %d) == %#016x; want %#016x", x, k, got, want)
			}
		}
	}
}

func TestReverse(t *testing.T) {
	// test each bit
	for i := uint(0); i < 64; i++ {
		testReverse(t, uint64(1)<<i, uint64(1)<<(63-i))
	}

	// test a few patterns
	for _, test := range []struct {
		x, r uint64
	}{
		{0, 0},
		{0x1, 0x8 << 60},
		{0x2, 0x4 << 60},
		{0x3, 0xc << 60},
		{0x4, 0x2 << 60},
		{0x5, 0xa << 60},
		{0x6, 0x6 << 60},
		{0x7, 0xe << 60},
		{0x8, 0x1 << 60},
		{0x9, 0x9 << 60},
		{0xa, 0x5 << 60},
		{0xb, 0xd << 60},
		{0xc, 0x3 << 60},
		{0xd, 0xb << 60},
		{0xe, 0x7 << 60},
		{0xf, 0xf << 60},
		{0x5686487, 0xe12616a000000000},
		{0x0123456789abcdef, 0xf7b3d591e6a2c480},
	} {
		testReverse(t, test.x, test.r)
		testReverse(t, test.r, test.x)
	}
}

func testReverse(t *testing.T, x64, want64 uint64) {
	x8 := uint8(x64)
	got8 := Reverse8(x8)
	want8 := uint8(want64 >> (64 - 8))
	if got8 != want8 {
		t.Fatalf("Reverse8(%#02x) == %#02x; want %#02x", x8, got8, want8)
	}

	x16 := uint16(x64)
	got16 := Reverse16(x16)
	want16 := uint16(want64 >> (64 - 16))
	if got16 != want16 {
		t.Fatalf("Reverse16(%#04x) == %#04x; want %#04x", x16, got16, want16)
	}

	x32 := uint32(x64)
	got32 := Reverse32(x32)
	want32 := uint32(want64 >> (64 - 32))
	if got32 != want32 {
		t.Fatalf("Reverse32(%#08x) == %#08x; want %#08x", x32, got32, want32)
	}
	if UintSize == 32 {
		x := uint(x32)
		got := Reverse(x)
		want := uint(want32)
		if got != want {
			t.Fatalf("Reverse(%#08x) == %#08x; want %#08x", x, got, want)
		}
	}

	got64 := Reverse64(x64)
	if got64 != want64 {
		t.Fatalf("Reverse64(%#016x) == %#016x; want %#016x", x64, got64, want64)
	}
	if UintSize == 64 {
		x := uint(x64)
		got := Reverse(x)
		want := uint(want64)
		if got != want {
			t.Fatalf("Reverse(%#08x) == %#016x; want %#016x", x, got, want)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(uint(i))
	}
}

func BenchmarkReverse8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse8(uint8(i))
	}
}

func BenchmarkReverse16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse16(uint16(i))
	}
}

func BenchmarkReverse32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse32(uint32(i))
	}
}

func BenchmarkReverse64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse64(uint64(i))
	}
}

func TestReverseBytes(t *testing.T) {
	for _, test := range []struct {
		x, r uint64
	}{
		{0, 0},
		{0x01, 0x01 << 56},
		{0x0123, 0x2301 << 48},
		{0x012345, 0x452301 << 40},
		{0x01234567, 0x67452301 << 32},
		{0x0123456789, 0x8967452301 << 24},
		{0x0123456789ab, 0xab8967452301 << 16},
		{0x0123456789abcd, 0xcdab8967452301 << 8},
		{0x0123456789abcdef, 0xefcdab8967452301 << 0},
	} {
		testReverseBytes(t, test.x, test.r)
		testReverseBytes(t, test.r, test.x)
	}
}

func testReverseBytes(t *testing.T, x64, want64 uint64) {
	x16 := uint16(x64)
	got16 := ReverseBytes16(x16)
	want16 := uint16(want64 >> (64 - 16))
	if got16 != want16 {
		t.Fatalf("ReverseBytes16(%#04x) == %#04x; want %#04x", x16, got16, want16)
	}

	x32 := uint32(x64)
	got32 := ReverseBytes32(x32)
	want32 := uint32(want64 >> (64 - 32))
	if got32 != want32 {
		t.Fatalf("ReverseBytes32(%#08x) == %#08x; want %#08x", x32, got32, want32)
	}
	if UintSize == 32 {
		x := uint(x32)
		got := ReverseBytes(x)
		want := uint(want32)
		if got != want {
			t.Fatalf("ReverseBytes(%#08x) == %#08x; want %#08x", x, got, want)
		}
	}

	got64 := ReverseBytes64(x64)
	if got64 != want64 {
		t.Fatalf("ReverseBytes64(%#016x) == %#016x; want %#016x", x64, got64, want64)
	}
	if UintSize == 64 {
		x := uint(x64)
		got := ReverseBytes(x)
		want := uint(want64)
		if got != want {
			t.Fatalf("ReverseBytes(%#016x) == %#016x; want %#016x", x, got, want)
		}
	}
}

func TestLen(t *testing.T) {
	for i := 0; i < 256; i++ {
		len := 8 - tab[i].nlz
		for k := 0; k < 64-8; k++ {
			x := uint64(i) << uint(k)
			want := 0
			if x != 0 {
				want = len + k
			}
			if x <= 1<<8-1 {
				got := Len8(uint8(x))
				if got != want {
					t.Fatalf("Len8(%#02x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<16-1 {
				got := Len16(uint16(x))
				if got != want {
					t.Fatalf("Len16(%#04x) == %d; want %d", x, got, want)
				}
			}

			if x <= 1<<32-1 {
				got := Len32(uint32(x))
				if got != want {
					t.Fatalf("Len32(%#08x) == %d; want %d", x, got, want)
				}
				if UintSize == 32 {
					got := Len(uint(x))
					if got != want {
						t.Fatalf("Len(%#08x) == %d; want %d", x, got, want)
					}
				}
			}

			if x <= 1<<64-1 {
				got := Len64(uint64(x))
				if got != want {
					t.Fatalf("Len64(%#016x) == %d; want %d", x, got, want)
				}
				if UintSize == 64 {
					got := Len(uint(x))
					if got != want {
						t.Fatalf("Len(%#016x) == %d; want %d", x, got, want)
					}
				}
			}
		}
	}
}

// ----------------------------------------------------------------------------
// Testing support

type entry = struct {
	nlz, ntz, pop int
}

// tab contains results for all uint8 values
var tab [256]entry

func init() {
	tab[0] = entry{8, 8, 0}
	for i := 1; i < len(tab); i++ {
		// nlz
		x := i // x != 0
		n := 0
		for x&0x80 == 0 {
			n++
			x <<= 1
		}
		tab[i].nlz = n

		// ntz
		x = i // x != 0
		n = 0
		for x&1 == 0 {
			n++
			x >>= 1
		}
		tab[i].ntz = n

		// pop
		x = i // x != 0
		n = 0
		for x != 0 {
			n += int(x & 1)
			x >>= 1
		}
		tab[i].pop = n
	}
}