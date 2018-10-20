package networking

import (
	"bytes"
	"encoding/binary"
	"math"
)

const (
	DeadStatus   = 0       // not alive, not paused, etc.
	Alive        = 1 << 0  // player is alive
	Paused       = 1 << 1  // player is paused
	Exploding    = 1 << 2  // currently blowing up
	Teleporting  = 1 << 3  // teleported recently
	FlagActive   = 1 << 4  // flag special powers active
	CrossingWall = 1 << 5  // tank crossing building wall
	Falling      = 1 << 6  // tank accel'd by gravity
	OnDriver     = 1 << 7  // tank is on a physics driver
	UserInputs   = 1 << 8  // user speed and angvel are sent
	JumpJets     = 1 << 9  // tank has jump jets on
	PlaySound    = 1 << 10 // play one or more sounds
)

const (
	smallScale     = 32766.0
	smallMaxDist   = 0.02 * smallScale
	smallMaxVel    = 0.01 * smallScale
	smallMaxAngVel = 0.001 * smallScale
)

type PlayerStateData struct {
	Position        Vector3F `json:"position"`
	Velocity        Vector3F `json:"velocity"`
	Azimuth         float32  `json:"azimuth"`
	AngularVelocity float32  `json:"angleVelocity"`

	physicsDriver int16
	userSpeed     float32
	userAngVel    float32
	jumpJetsScale float32
	sounds        uint8
}

func UnpackPlayerState(buf *bytes.Buffer, code uint16) (packet PlayerStateData) {
	var inOrder uint32
	var inStatus uint16

	binary.Read(buf, binary.BigEndian, &inOrder)
	binary.Read(buf, binary.BigEndian, &inStatus)

	if code == codeFromChars(MsgPlayerUpdate) {
		packet.Position = UnpackVector(buf)
		packet.Velocity = UnpackVector(buf)
		packet.Azimuth = UnpackFloat(buf)
		packet.AngularVelocity = UnpackFloat(buf)
	} else {
		var pos, vel [3]int16
		var azi, angVel int16

		binary.Read(buf, binary.BigEndian, &pos[0])
		binary.Read(buf, binary.BigEndian, &pos[1])
		binary.Read(buf, binary.BigEndian, &pos[2])
		binary.Read(buf, binary.BigEndian, &vel[0])
		binary.Read(buf, binary.BigEndian, &vel[1])
		binary.Read(buf, binary.BigEndian, &vel[2])
		binary.Read(buf, binary.BigEndian, &azi)
		binary.Read(buf, binary.BigEndian, &angVel)

		for i := 0; i < 3; i++ {
			packet.Position[i] = (float32(pos[i]) * smallMaxDist) / smallScale
			packet.Velocity[i] = (float32(vel[i]) * smallMaxVel) / smallScale
		}

		packet.Azimuth = (float32(azi) * math.Pi) / smallScale
		packet.AngularVelocity = (float32(angVel) * smallMaxAngVel) / smallScale
	}

	if inStatus&JumpJets != 0 {
		var jumpJetsValue uint16
		binary.Read(buf, binary.BigEndian, &jumpJetsValue)

		packet.jumpJetsScale = float32(jumpJetsValue) / smallScale
	} else {
		packet.jumpJetsScale = 0.0
	}

	if inStatus&OnDriver != 0 {
		var inPhyDrv uint32
		binary.Read(buf, binary.BigEndian, &inPhyDrv)
		packet.physicsDriver = int16(inPhyDrv)
	} else {
		packet.physicsDriver = -1
	}

	if inStatus&UserInputs != 0 {
		var speed, angVel uint16
		binary.Read(buf, binary.BigEndian, &speed)
		binary.Read(buf, binary.BigEndian, &angVel)

		packet.userSpeed = (float32(speed) * smallMaxVel) / smallScale
		packet.userAngVel = (float32(angVel) * smallMaxAngVel) / smallScale
	} else {
		packet.userSpeed = 0
		packet.userAngVel = 0
	}

	if inStatus&PlaySound != 0 {
		binary.Read(buf, binary.BigEndian, &packet.sounds)
	} else {
		packet.sounds = 0
	}

	return
}
