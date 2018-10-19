package networking

import (
	"bytes"
	"encoding/binary"
)

const (
	FlagNoExist  = iota // the flag is not present in the world
	FlagOnGround        // the flag is sitting on the ground and can be picked up
	FlagOnTank          // the flag is being carried by a tank
	FlagInAir           // the flag is falling through the air
	FlagComing          // the flag is entering the world
	FlagGoing           // the flag is leaving the world
)

const (
	FlagNormal   = iota // permanent flag
	FlagUnstable        // disappears after use
	FlagSticky          // can't be dropped normally
)

type MsgFlagUpdatePacket struct {
	Type  string
	Flags []FlagData
}

type FlagData struct {
	Index           uint16         `json:"index"`
	Abbv            string         `json:"abbv"`
	Status          uint16         `json:"status"`
	Endurance       uint16         `json:"endurance"`
	Owner           uint8          `json:"owner"`
	Position        Vector3F `json:"position"`
	LaunchPosition  Vector3F `json:"launchPos"`
	LandingPosition Vector3F `json:"landingPos"`
	FlightTime      float32        `json:"flightTime"`
	FlightEnd       float32        `json:"flightEnd"`
	InitialVelocity float32        `json:"initialVelocity"`
}

func (m *MsgFlagUpdatePacket) Unpack(data []byte) (unpacked MsgFlagUpdatePacket) {
	buf := bytes.NewBuffer(data)

	unpacked.Type = "MsgFlagUpdate"

	var count uint16
	binary.Read(buf, binary.BigEndian, &count)

	var i uint16
	for i = 0; i < count; i++ {
		var flag FlagData

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

		unpacked.Flags = append(unpacked.Flags, flag)
	}

	return
}
