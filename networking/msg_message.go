package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgMessagePacket struct {
	Type         string `json:"type"`
	PlayerFromID uint8  `json:"playerFromID"`
	PlayerToID   uint8  `json:"playerToID"`
	Message      string `json:"message"`
}

func (m *MsgMessagePacket) Unpack(buf *bytes.Buffer) (packet MsgMessagePacket) {
	packet.Type = "MsgMessage"

	binary.Read(buf, binary.BigEndian, &packet.PlayerFromID)
	binary.Read(buf, binary.BigEndian, &packet.PlayerToID)

	packet.Message = UnpackString(buf, buf.Len())

	return
}
