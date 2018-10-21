package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgTransferFlagPacket struct {
	Type string   `json:"type"`
	From uint8    `json:"from"`
	To   uint8    `json:"to"`
	Flag FlagData `json:"flag"`
}

func (m *MsgTransferFlagPacket) Unpack(buf *bytes.Buffer) (packet MsgTransferFlagPacket) {
	packet.Type = "MsgTransferFlag"

	binary.Read(buf, binary.BigEndian, &packet.From)
	binary.Read(buf, binary.BigEndian, &packet.To)

	packet.Flag = UnpackFlag(buf)

	return
}
