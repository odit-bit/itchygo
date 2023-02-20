package itchygo

import (
	"encoding/binary"
	"errors"
)

var ErrIntegerWrongBinaryLength = errors.New("wrong binary length")

// will decode to int using specified big-endian binary-order uint16/32/64  based on byte length
func Integer(b []byte) (int, error) {
	//fmt.Println("prcosessing ", b)
	var result int
	length := len(b)

	//uint16
	if length == 2 {
		return u16(b), nil
	}
	//uint32
	if length == 4 {
		return u32(b), nil
	}
	//uint64
	if length == 8 {
		return u64(b), nil
	}

	return result, errors.New("wrong binary length")
}

// read b using specified big-endian uint16
func u16(b []byte) int {
	result := int(binary.BigEndian.Uint16(b))
	return result
}

// read b using specified big-endian uint32
func u32(b []byte) int {
	result := int(binary.BigEndian.Uint32(b))
	return result
}

// read b using specified big-endian uint64
func u64(b []byte) int {
	result := int(binary.BigEndian.Uint64(b))
	return result
}

func AlphaString(b []byte) string {
	return string(b)
}

// every itch message contain type as Integer,Alpha, null-terminated.
// those type has different length byte such as 4 , 8
// this function used to translate to those types

type U8 []byte
type U16 []byte
type U32 []byte
type U64 []byte

func Len1(i int) U8 {
	return []byte{uint8(i)}
}

// when to translate itch Integer with 2 length need uint16
func Len2(i int) U16 {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(i))
	return b
}

// when to translate itch Integer with 4 length need uint32
func Len4(i int) U32 {
	b := make([]byte, 4)
	binary.BigEndian.PutUint32(b, uint32(i))
	return b
}

// when to translate itch Integer with 8 length need uint64
func Len8(i int) U64 {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(i))
	return b
}

type Alpha []byte

// convert s into byte, if b size not enough it will panic
func AlphaByte(s string, size int) ([]byte, error) {
	b := make([]byte, size)
	for i := range s {
		b[i] = s[i]
	}
	return b, nil
}
