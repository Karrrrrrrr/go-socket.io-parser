package socket_io_parser

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type Encoder struct {
}
type PacketType int

const (
	CONNECT PacketType = iota
	DISCONNECT
	EVENT
	ACK
	ConnectError
	BinaryEvent
	BinaryAck
	MaxPacketType
)

type Packet struct {
	Type        PacketType
	Nsp         string
	Data        any
	Id          int
	Attachments int
}

func (e *Encoder) Encode(obj Packet) ([]byte, error) {
	if obj.Type == EVENT || obj.Type == ACK {

	}
	return e.encodeAsString(obj)

	// if (obj.type === PacketType.EVENT || obj.type === PacketType.ACK) {
	//  if (hasBinary(obj)) {
	//    return this.encodeAsBinary({
	//      type:
	//        obj.type === PacketType.EVENT
	//          ? PacketType.BINARY_EVENT
	//          : PacketType.BINARY_ACK,
	//      nsp: obj.nsp,
	//      data: obj.data,
	//      id: obj.id,
	//    });
	//  }
	//}
}

func (e *Encoder) encodeAsString(obj Packet) ([]byte, error) {
	var buf = bytes.NewBufferString("")
	buf.WriteString(strconv.Itoa(int(obj.Type)))

	if obj.Type == BinaryEvent || obj.Type == BinaryAck {
		buf.WriteString(strconv.Itoa(obj.Attachments) + "-")
	}
	if obj.Nsp != "" && "/" != obj.Nsp {
		buf.WriteString(obj.Nsp + ",")
	}
	if obj.Id >= 0 {
		buf.WriteString(strconv.Itoa(obj.Id))
	}
	if nil != obj.Data {
		marshal, err := json.Marshal(obj.Data)
		if err != nil {
			return nil, err
		}
		buf.Write(marshal)
	}
	return buf.Bytes(), nil
}
func (e *Encoder) encodeAsBinary() {

}

/*
export function deconstructPacket(packet) {
  const buffers = [];
  const packetData = packet.data;
  const pack = packet;
  pack.data = _deconstructPacket(packetData, buffers);
  pack.attachments = buffers.length; // number of binary 'attachments'
  return { packet: pack, buffers: buffers };
}


function _deconstructPacket(data, buffers) {
  if (!data) return data;

  if (isBinary(data)) {
    const placeholder = { _placeholder: true, num: buffers.length };
    buffers.push(data);
    return placeholder;
  } else if (Array.isArray(data)) {
    const newData = new Array(data.length);
    for (let i = 0; i < data.length; i++) {
      newData[i] = _deconstructPacket(data[i], buffers);
    }
    return newData;
  } else if (typeof data === "object" && !(data instanceof Date)) {
    const newData = {};
    for (const key in data) {
      if (Object.prototype.hasOwnProperty.call(data, key)) {
        newData[key] = _deconstructPacket(data[key], buffers);
      }
    }
    return newData;
  }
  return data;
}

*/

func deconstructPacket(packet Packet) {

}

// todo
func isBinary(obj any) bool {
	return false
}

/*

export function isBinary(obj: any) {
  return (
    (withNativeArrayBuffer && (obj instanceof ArrayBuffer || isView(obj))) ||
    (withNativeBlob && obj instanceof Blob) ||
    (withNativeFile && obj instanceof File)
  );
}

export function hasBinary(obj: any, toJSON?: boolean) {
  if (!obj || typeof obj !== "object") {
    return false;
  }

  if (Array.isArray(obj)) {
    for (let i = 0, l = obj.length; i < l; i++) {
      if (hasBinary(obj[i])) {
        return true;
      }
    }
    return false;
  }

  if (isBinary(obj)) {
    return true;
  }

  if (
    obj.toJSON &&
    typeof obj.toJSON === "function" &&
    arguments.length === 1
  ) {
    return hasBinary(obj.toJSON(), true);
  }

  for (const key in obj) {
    if (Object.prototype.hasOwnProperty.call(obj, key) && hasBinary(obj[key])) {
      return true;
    }
  }

  return false;
}
*/
