package networking

import (
	"bytes"
)

type FiringInfoData struct {
	TimeSent float32  `json:"timeSent"`
	Shot     ShotData `json:"shot"`
	Flag     FlagData `json:"flag"`
	Lifetime float32  `json:"lifetime"`
}

func UnpackFiringInfo(buf *bytes.Buffer) (packet FiringInfoData) {
	packet.TimeSent = UnpackFloat(buf)
	packet.Shot = UnpackShot(buf)
	packet.Flag = UnpackFlag(buf)
	packet.Lifetime = UnpackFloat(buf)

	return
}
