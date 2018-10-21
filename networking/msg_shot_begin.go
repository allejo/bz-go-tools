package networking

import "bytes"

type MsgShotBeginPacket struct {
	Type       string         `json:"type"`
	FiringInfo FiringInfoData `json:"firingInfo"`
}

func (m *MsgShotBeginPacket) Unpack(buf *bytes.Buffer) (packet MsgShotBeginPacket) {
	packet.Type = "MsgShotBegin"
	packet.FiringInfo = UnpackFiringInfo(buf)

	return
}
