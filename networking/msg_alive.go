package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgAlivePacket struct {
	Type     string   `json:"type"`
	Player   uint8    `json:"playerID"`
	Position Vector3F `json:"position"`
	Azimuth  float32  `json:"azimuth"`
}

func (m *MsgAlivePacket) Unpack(buf *bytes.Buffer) (packet MsgAlivePacket) {
	packet.Type = "MsgAlive"

	binary.Read(buf, binary.BigEndian, &packet.Player)

	packet.Position = UnpackVector(buf)
	packet.Azimuth = UnpackFloat(buf)

	return
}
