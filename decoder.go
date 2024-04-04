package socket_io_parser

import (
	"encoding/json"
	"errors"
	"strconv"
)

type Decoder struct {
}

func (d Decoder) Decode(data []any) (pack *Packet, err error) {
	if len(data) == 0 {
		return nil, errors.New("error")
	}
	if s, ok := data[0].(string); len(data) == 1 && ok {
		return d.decodeAsString(s)
	}
	return nil, errors.New("unknown error")

	//for i := range data {
	//
	//}
}

func (d Decoder) decodeAsString(s string) (*Packet, error) {
	var err error
	var i = 0
	var p = Packet{}
	p.Type = PacketType(s[0] - '0')
	if p.Type < 0 || p.Type >= MaxPacketType {
		return &p, errors.New("error")
	}
	// look up attachments if type binary
	if p.Type == BinaryEvent || p.Type == BinaryAck {
		var start = i + 1
		for {
			i++
			if i == len(s) || s[i] == '-' {
				break
			}
		}
		var buf = s[start:i]
		p.Attachments, err = strconv.Atoi(buf)
		if err != nil || s[i] != '-' {
			return nil, errors.New("illegal attachments")
		}

	}
	// look up namespace (if any)
	p.Nsp = "/"
	if s[i+1] == '/' {
		var start = i + 1
		for {
			i++
			if i == len(s) || s[i] == ',' {
				break
			}
		}
		p.Nsp = s[start:i]
	}
	// look up id
	if start := i + 1; start < len(s) && isDigit(s[start]) {
		for {
			i++
			if i >= len(s) || !isDigit(s[i]) {
				i--
				break
			}
			// ?
			if i == len(s) {
				break
			}
		}
		p.Id, err = strconv.Atoi(s[start : i+1])
	}
	// look up json data
	i++
	if i < len(s) {
		var payload = s[i:]
		var mp any
		var err = json.Unmarshal([]byte(payload), &mp)
		if err != nil {
			return nil, errors.New("invalid payload")
		}
		p.Data = json.RawMessage(payload)
	}
	return &p, nil
}

func isDigit(b byte) bool {
	return b <= '9' && b >= '0'
}

/*
 let packet;
    if (typeof obj === "string") {
      if (this.reconstructor) {
        throw new Error("got plaintext data when reconstructing a packet");
      }
      packet = this.decodeString(obj);
      const isBinaryEvent = packet.type === PacketType.BINARY_EVENT;
      if (isBinaryEvent || packet.type === PacketType.BINARY_ACK) {
        packet.type = isBinaryEvent ? PacketType.EVENT : PacketType.ACK;
        // binary packet's json
        this.reconstructor = new BinaryReconstructor(packet);

        // no attachments, labeled binary but no binary data to follow
        if (packet.attachments === 0) {
          super.emitReserved("decoded", packet);
        }
      } else {
        // non-binary full packet
        super.emitReserved("decoded", packet);
      }
    } else if (isBinary(obj) || obj.base64) {
      // raw binary data
      if (!this.reconstructor) {
        throw new Error("got binary data when not reconstructing a packet");
      } else {
        packet = this.reconstructor.takeBinaryData(obj);
        if (packet) {
          // received final buffer
          this.reconstructor = null;
          super.emitReserved("decoded", packet);
        }
      }
    } else {
      throw new Error("Unknown type: " + obj);
    }
  }
*/
