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

func (m *MsgFlagUpdatePacket) Unpack(buf *bytes.Buffer) (packet MsgFlagUpdatePacket) {
	packet.Type = "MsgFlagUpdate"

	var count uint16
	binary.Read(buf, binary.BigEndian, &count)

	var i uint16
	for i = 0; i < count; i++ {
		flag := UnpackFlag(buf)
		packet.Flags = append(packet.Flags, flag)
	}

	return
}
