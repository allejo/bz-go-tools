package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgGMUpdatePacket struct {
	Type   string   `json:"type"`
	Target uint8    `json:"target"`
	Shot   ShotData `json:"shot"`
}

func (m *MsgGMUpdatePacket) Unpack(buf *bytes.Buffer) (packet MsgGMUpdatePacket) {
	packet.Type = "MsgGMUpdate"
	packet.Shot = UnpackShot(buf)

	binary.Read(buf, binary.BigEndian, &packet.Target)

	return
}
