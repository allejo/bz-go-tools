package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgTeleportPacket struct {
	Type     string `json:"type"`
	PlayerID uint8  `json:"playerID"`
	From     uint16 `json:"from"`
	To       uint16 `json:"to"`
}

func (m *MsgTeleportPacket) Unpack(buf *bytes.Buffer) (packet MsgTeleportPacket) {
	packet.Type = "MsgTeleport"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)
	binary.Read(buf, binary.BigEndian, &packet.From)
	binary.Read(buf, binary.BigEndian, &packet.To)

	return
}
