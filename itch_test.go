package itchygo

import (
	"bytes"
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

var testMap map[int][]byte = map[int][]byte{
	4009: func() []byte {
		msg := []byte{}
		msg = binary.BigEndian.AppendUint16(msg, uint16(4009))
		return msg
	}(),
	4200000009: func() []byte {
		msg := []byte{}
		msg = binary.BigEndian.AppendUint32(msg, uint32(4200000009))
		return msg
	}(),
	420000000000000009: func() []byte {
		msg := []byte{}
		msg = binary.BigEndian.AppendUint64(msg, uint64(420000000000000009))
		return msg
	}(),
}

func Test_Int(t *testing.T) {

	r := require.New(t)

	for k, v := range testMap {
		actual, err := Integer(v)
		r.NoError(err, err)
		r.Equal(k, actual, "")
	}

	//test error
	i, err := Integer([]byte{1})
	r.Error(err, "should error")
	r.Equal(err, ErrIntegerWrongBinaryLength, "wrong binary length")
	r.Equal(0, i, "should same")
}

// func Test_Alpha(t *testing.T) {
// 	r := require.New(t)

// 	msg := []byte("odit-bit finance")
// 	AlphaString(msg)

// }

var msg []byte = []byte("odit-bit finance")

func BenchmarkAlpha(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AlphaString(msg)
	}
}

func generateItchMsg() []byte {
	/*{
	"Type":"E", len 1
	"Timestamp":90963000, len 4
	"OrderNumber":202103060000028634, len 8
	"ExecutedQuantity":1, len 8
	"MatchNumber":202103060000016258,
	"TradeIndicator":"R", len 1
	"BuyParticipantId":44, len 4
	"SellParticipantId":44}, len 4
	*/
	b := make([]byte, 40)

	//type
	//b[0]
	str := uint8('E')
	b[0] = str

	//timestamp
	//b[1:] length = 4
	// t, _ := time.Parse(time.DateTime, "2023-02-16 00:00:00")
	// timestamp := time.Since(t)
	binary.BigEndian.PutUint32(b[1:5], uint32(90963000))

	//order Number
	//b[5:] length = 8
	orderNum := uint64(202103060000028634)
	//binary.PutUvarint(b[5:13], orderNum)
	binary.BigEndian.PutUint64(b[5:13], orderNum)

	//Executed quantity
	//b[13:] length = 8
	exeQuant := 1
	binary.BigEndian.PutUint64(b[13:21], uint64(exeQuant))

	//Match Number
	//b[21:] length 8
	matchNum := 202103060000016258
	binary.BigEndian.PutUint64(b[21:29], uint64(matchNum))

	//Trade indicator
	//b[29] length = 1
	tradeInd := 'R'
	b[29] = uint8(tradeInd)

	//buy participant id
	//b[30:] length 4
	buyPartId := 44
	binary.BigEndian.PutUint32(b[30:34], uint32(buyPartId))

	//sell participant id
	//b[34:] length 4
	sellPartId := 44
	binary.BigEndian.PutUint32(b[34:38], uint32(sellPartId))

	//buy domicile
	//b[38] length 1
	buyDom := 'I'
	b[38] = uint8(buyDom)

	//sell domicile
	sellDom := 'A'
	b[39] = uint8(sellDom)

	// c := [40]byte{}
	// for i := 0; i < 40; i++ {
	// 	c[i] = b[i]
	// }

	// fmt.Println("byte order", c)
	// fmt.Println("Length", len(c))
	return b
}

func Test_DecodeStream(t *testing.T) {
	r := require.New(t)
	buf := bytes.NewBuffer(generateItchMsg())
	//read types
	b := buf.Next(1)
	typ := string(b)
	r.Equal("E", typ, "")

	//read next, timeStamp
	b = buf.Next(4)
	timestamp, err := Integer(b)
	r.NoError(err, err)
	r.Equal(90963000, timestamp)

	//read next, order number
	b = buf.Next(8)
	orderNum, err := Integer(b)
	r.NoError(err, err)
	r.Equal(202103060000028634, orderNum, "")

	//read next, Executed quantity
	b = buf.Next(8)
	exeQuant, err := Integer(b)
	r.NoError(err, err)
	r.Equal(1, exeQuant, "")

	//read next, MatchNum
	b = buf.Next(8)
	matchNum, err := Integer(b)
	r.NoError(err, err)
	r.Equal(202103060000016258, matchNum, "match num")

	//read next, TradeInd
	b = buf.Next(1)
	tradeInd := AlphaString(b)
	r.NoError(err, err)
	r.Equal("R", tradeInd, "tradeInd")

	//read next, BuyParty
	b = buf.Next(4)
	bp, err := Integer(b)
	r.NoError(err, err)
	r.Equal(44, bp, "buy participant")

	//read next, sellParty
	b = buf.Next(4)
	sp, err := Integer(b)
	r.NoError(err, err)
	r.Equal(44, sp, "sell participant")

	//read next, buy domicile
	b = buf.Next(1)
	_ = b
	//read next, sell domicile
	b = buf.Next(1)
	_ = b

	//reader should exhausted
	r.Equal(0, buf.Len(), "reader length")
}
