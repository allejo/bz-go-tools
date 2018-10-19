package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgDropFlagPacket struct {
	Type     string   `json:"type"`
	PlayerID uint8    `json:"playerID"`
	Flag     FlagData `json:"flag"`
}

func (m *MsgDropFlagPacket) Unpack(buf *bytes.Buffer) (packet MsgDropFlagPacket) {
	packet.Type = "MsgDropFlag"

	binary.Read(buf, binary.BigEndian, &packet.PlayerID)

	packet.Flag = UnpackFlag(buf)

	return
}
