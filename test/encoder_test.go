package test

import (
	"encoding/json"
	"fmt"
	socket_io_parser "github.com/karrrrrrrr/go-socket.io-parser"
	"testing"
)

func TestEncoder(t *testing.T) {
	encoder := socket_io_parser.Encoder{}
	encode, _ := encoder.Encode(socket_io_parser.Packet{
		Type:        2,
		Nsp:         "",
		Data:        []any{"hello world", json.RawMessage(`{"id":1, "name":"zhang san"}`)},
		Id:          -1,
		Attachments: 0,
	})
	fmt.Println(string(encode))

	decoder := socket_io_parser.Decoder{}
	pack, err := decoder.Decode([]any{string(encode)})
	if nd, ok := pack.Data.(json.RawMessage); ok {
		fmt.Println("nd=", string(nd))
	}
	fmt.Printf("%+v %+v\n", pack, err)
}
