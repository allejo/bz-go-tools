package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgNewRabbitPacket struct {
	Type     string `json:"type"`
	PlayerID uint8  `json:"playerID"`
	Paused   uint8  `json:"paused"`
}

func (m *MsgNewRabbitPacket) Unpack(buf *bytes.Buffer) (packet MsgNewRabbitPacket) {
	packet.Type = "MsgNewRabbit"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)
	binary.Read(buf, binary.BigEndian, &packet.Paused)

	return
}
