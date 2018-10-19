package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgCaptureFlagPacket struct {
	Type string `json:"type"`
	Team uint16 `json:"team"`
}

func (m *MsgCaptureFlagPacket) Unpack(buf *bytes.Buffer) (packet MsgCaptureFlagPacket) {
	packet.Type = "MsgCaptureFlag"

	binary.Read(buf, binary.BigEndian, &packet.Team)

	return
}
