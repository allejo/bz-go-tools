package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgShotEndPacket struct {
	Type     string `json:"type"`
	PlayerID uint8  `json:"playerID"`
	ShotID   uint16 `json:"shotID"`
	Reason   int16  `json:"reason"`
}

func (m *MsgShotEndPacket) Unpack(buf *bytes.Buffer) (packet MsgShotEndPacket) {
	packet.Type = "MsgShotEnd"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)
	binary.Read(buf, binary.BigEndian, &packet.ShotID)
	binary.Read(buf, binary.BigEndian, &packet.Reason)

	return
}
