package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgPausePacket struct {
	Type     string `json:"type"`
	PlayerID uint8  `json:"playerID"`
	Paused   uint8  `json:"paused"`
}

func (m *MsgPausePacket) Unpack(buf *bytes.Buffer) (packet MsgPausePacket) {
	packet.Type = "MsgPause"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)
	binary.Read(buf, binary.BigEndian, &packet.Paused)

	return
}
