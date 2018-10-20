package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgPlayerUpdatePacket struct {
	Type      string          `json:"type"`
	Timestamp float32         `json:"timestamp"`
	PlayerID  uint8           `json:"playerID"`
	State     PlayerStateData `json:"state"`
}

func (m *MsgPlayerUpdatePacket) Unpack(buf *bytes.Buffer, code uint16) (packet MsgPlayerUpdatePacket) {
	packet.Type = "MsgPlayerUpdate"

	binary.Read(buf, binary.BigEndian, &packet.Timestamp)
	binary.Read(buf, binary.BigEndian, &packet.PlayerID)

	packet.State = UnpackPlayerState(buf, code)

	return
}
