package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgScoreOverPacket struct {
	Type     string `json:"type"`
	PlayerID uint8  `json:"playerID"`
	Team     uint16 `json:"team"`
}

func (m *MsgScoreOverPacket) Unpack(buf *bytes.Buffer) (packet MsgScoreOverPacket) {
	packet.Type = "MsgScoreOver"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)
	binary.Read(buf, binary.BigEndian, &packet.Team)

	return
}
