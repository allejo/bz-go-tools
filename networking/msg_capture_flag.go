package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgCaptureFlagPacket struct {
	Type     string `json:"type"`
	PlayerID uint8  `json:"playerID"`
	FlagID   uint16 `json:"flagID"`
	Team     uint16 `json:"team"`
}

func (m *MsgCaptureFlagPacket) Unpack(buf *bytes.Buffer) (packet MsgCaptureFlagPacket) {
	packet.Type = "MsgCaptureFlag"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)
	binary.Read(buf, binary.BigEndian, &packet.FlagID)
	binary.Read(buf, binary.BigEndian, &packet.Team)

	return
}
