package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgTimeUpdatePacket struct {
	Type     string `json:"type"`
	TimeLeft int32  `json:"timeLeft"`
}

func (m *MsgTimeUpdatePacket) Unpack(buf *bytes.Buffer) (packet MsgTimeUpdatePacket) {
	packet.Type = "MsgTimeUpdate"

	binary.Read(buf, binary.BigEndian, &packet.TimeLeft)

	return
}
