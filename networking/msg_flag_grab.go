package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgFlagGrabPacket struct {
	Type     string   `json:"type"`
	PlayerID uint8    `json:"playerID"`
	Flag     FlagData `json:"flag"`
}

func (m *MsgFlagGrabPacket) Unpack(buf *bytes.Buffer) (packet MsgFlagGrabPacket) {
	packet.Type = "MsgFlagGrab"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)

	packet.Flag = UnpackFlag(buf)

	return
}
