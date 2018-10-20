package networking

import (
	"bytes"
	"encoding/binary"
)

type ShotData struct {
	PlayerID  uint8    `json:"playerID"`
	ShotID    uint16   `json:"shotID"`
	Position  Vector3F `json:"position"`
	Velocity  Vector3F `json:"velocity"`
	DeltaTime float32  `json:"deltaTime"`
	Team      uint16   `json:"team"`
}

func UnpackShot(buf *bytes.Buffer) (shot ShotData) {
	binary.Read(buf, binary.BigEndian, &shot.PlayerID)
	binary.Read(buf, binary.BigEndian, &shot.ShotID)

	shot.Position = UnpackVector(buf)
	shot.Velocity = UnpackVector(buf)
	shot.DeltaTime = UnpackFloat(buf)

	binary.Read(buf, binary.BigEndian, &shot.Team)

	return
}
