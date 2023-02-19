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

	// if length == 1 {
	// 	return int(b[0]), nil
	// }

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

// convert s int b, if b size not enough it will panic
func Alpha(s string, size int) ([]byte, error) {
	b := make([]byte, size)
	for i := range s {
		b[i] = s[i]
	}
	return b, nil
}

func AlphaString(b []byte) string {
	// var s strings.Builder

	// s.Grow(len(b))
	// for _, v := range b {
	// 	s.WriteByte(v)
	// }
	// return s.String()
	return string(b)
}
