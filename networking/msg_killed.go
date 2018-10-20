package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgKilledPacket struct {
	Type            string   `json:"type"`
	VictimID        uint8    `json:"victimID"`
	KillerID        uint8    `json:"killerID"`
	Reason          uint16   `json:"reason"`
	ShotID          uint16   `json:"shotID"`
	Flag            FlagData `json:"flag"`
	PhysicsDriverID uint32   `json:"physicsDriverID"`
}

func (m *MsgKilledPacket) Unpack(buf *bytes.Buffer) (packet MsgKilledPacket) {
	packet.Type = "MsgKilled"

	binary.Read(buf, binary.BigEndian, &packet.VictimID)
	binary.Read(buf, binary.BigEndian, &packet.KillerID)
	binary.Read(buf, binary.BigEndian, &packet.Reason)
	binary.Read(buf, binary.BigEndian, &packet.ShotID)

	packet.Flag = UnpackFlag(buf)

	// Physics Death are a network message for legacy reasons
	if packet.Reason == codeFromChars("pd") {
		binary.Read(buf, binary.BigEndian, &packet.PhysicsDriverID)
	}

	return
}
