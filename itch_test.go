package itchygo

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// TRANSLATE TO and FROM ITCH BYTE

func Test_AlphaByte(t *testing.T) {
	r := require.New(t)
	s := "S"
	a, err := AlphaByte(s, len(s))
	r.NoError(err, "no error")
	actual := AlphaString(a)
	r.Equal(s, actual, "AlphaByte")
}

func Test_len2(t *testing.T) {
	r := require.New(t)
	i := 65535
	b := Len2(i)
	actual, err := Integer(b)
	r.NoError(err, "no error")
	r.Equal(i, actual, "should same")
}

func Test_len4(t *testing.T) {
	r := require.New(t)
	i := 4200000000
	b := Len4(i)
	actual, err := Integer(b)
	r.NoError(err, "no error")
	r.Equal(i, actual, "should same")
}

func Test_len8(t *testing.T) {
	r := require.New(t)
	i := 8999999999999999999
	b := Len8(i)
	actual, err := Integer(b)
	r.NoError(err, "no error")
	r.Equal(i, actual, "should same")
}

// Benchmark
var msg []byte = []byte("odit-bit finance")

func BenchmarkAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaString(msg)
	}
}

var l4 int = 6553
var l8 int = 12345678910

func Benchmark_l4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b U32
		b = Len4(l4)
		_ = b
	}
}

func Benchmark_l8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		var b U64
		b = Len8(l8)
		_ = b
	}
}
