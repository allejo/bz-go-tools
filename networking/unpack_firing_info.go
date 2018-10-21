package networking

import (
	"bytes"
)

type FiringInfoData struct {
	TimeSent float32  `json:"timeSent"`
	Shot     ShotData `json:"shot"`
	Flag     string   `json:"flag"`
	Lifetime float32  `json:"lifetime"`
}

func UnpackFiringInfo(buf *bytes.Buffer) (packet FiringInfoData) {
	packet.TimeSent = UnpackFloat(buf)
	packet.Shot = UnpackShot(buf)
	packet.Flag = UnpackString(buf, 2)
	packet.Lifetime = UnpackFloat(buf)

	return
}
