package networking

import (
	"bytes"
	"encoding/binary"
)

type FlagData struct {
	Index           uint16   `json:"index"`
	Abbv            string   `json:"abbv"`
	Status          uint16   `json:"status"`
	Endurance       uint16   `json:"endurance"`
	Owner           uint8    `json:"owner"`
	Position        Vector3F `json:"position"`
	LaunchPosition  Vector3F `json:"launchPos"`
	LandingPosition Vector3F `json:"landingPos"`
	FlightTime      float32  `json:"flightTime"`
	FlightEnd       float32  `json:"flightEnd"`
	InitialVelocity float32  `json:"initialVelocity"`
}

func UnpackFlag(buf *bytes.Buffer) (flag FlagData) {
	binary.Read(buf, binary.BigEndian, &flag.Index)

	flag.Abbv = UnpackString(buf, 2)

	binary.Read(buf, binary.BigEndian, &flag.Status)
	binary.Read(buf, binary.BigEndian, &flag.Endurance)
	binary.Read(buf, binary.BigEndian, &flag.Owner)

	flag.Position = UnpackVector(buf)
	flag.LaunchPosition = UnpackVector(buf)
	flag.LandingPosition = UnpackVector(buf)
	flag.FlightTime = UnpackFloat(buf)
	flag.FlightEnd = UnpackFloat(buf)
	flag.InitialVelocity = UnpackFloat(buf)

	return
}
