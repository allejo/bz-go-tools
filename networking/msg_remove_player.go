package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgRemovePlayerPacket struct {
	Type     string `json:"type"`
	PlayerID uint8  `json:"playerID"`
}

func (m *MsgRemovePlayerPacket) Unpack(buf *bytes.Buffer) (packet MsgRemovePlayerPacket) {
	packet.Type = "MsgPlayerRemove"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)

	return
}
