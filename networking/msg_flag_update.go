package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgFlagUpdatePacket struct {
	Type  string     `json:"type"`
	Flags []FlagData `json:"flags"`
}

func (m *MsgFlagUpdatePacket) Unpack(buf *bytes.Buffer) (packet MsgFlagUpdatePacket) {
	packet.Type = "MsgFlagUpdate"

	var count uint16
	binary.Read(buf, binary.BigEndian, &count)

	var i uint16
	for i = 0; i < count; i++ {
		flag := UnpackFlag(buf)
		packet.Flags = append(packet.Flags, flag)
	}

	return
}
