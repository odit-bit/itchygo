package itch

import "testing"

var data = []struct {
	in  []byte
	out Message
}{
	{in: []byte{'T', 0, 0, 98, 112}, out: ItchTime{Second: 25200}},
	{
		in: []byte{'A', 29, 205, 101, 0, 0, 0, 0, 0, 0, 0, 0, 99, 'B', 0, 0, 0, 0, 0, 0, 0, 100, 0, 0, 6, 236, 0, 0, 78, 32},
		out: AddOrder{
			Timestamp: 500000,
			OrderNum:  99,
			OrderVerb: 'B',
			Quantity:  100,
			OrderBook: 1772,
			Price:     20000,
		},
	},
}

// func addOneStream() []byte {
// 	stream := []byte{}

// 	for _, v := range data {
// 		stream = append(stream, v.in...)
// 	}
// 	return stream
// }

var serializer = NewSerializer()
var result Message

func Benchmark_Serialize_timestamp(b *testing.B) {
	var it ItchTime
	fn := serializer[TIMESTAMP_MESSAGE]
	for i := 0; i < b.N; i++ {
		err := fn(data[0].in, &it)
		if err != nil {
			b.Fatal(err)
		}
	}
	result = it
}

func Benchmark_Serialize_AddOrder(b *testing.B) {
	var ao AddOrder
	fn := serializer[ADD_ORDER_MESSAGE]
	for i := 0; i < b.N; i++ {
		err := fn(data[1].in, &ao)
		if err != nil {
			b.Fatal(err)
		}
	}
	result = ao
}
