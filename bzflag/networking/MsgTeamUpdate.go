package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgTeamUpdatePacket struct {
	Type  string     `json:"type"`
	Teams []TeamData `json:"teams"`
}

type TeamData struct {
	Team   uint16 `json:"team"`
	Size   uint16 `json:"size"`
	Wins   uint16 `json:"wins"`
	Losses uint16 `json:"losses"`
}

func (m *MsgTeamUpdatePacket) Unpack(data []byte) (unpacked MsgTeamUpdatePacket) {
	buf := bytes.NewBuffer(data)

	unpacked.Type = "MsgTeamUpdate"

	var count uint8
	binary.Read(buf, binary.BigEndian, &count)

	var i uint8
	for i = 0; i < count; i++ {
		var data TeamData

		binary.Read(buf, binary.BigEndian, &data.Team)
		binary.Read(buf, binary.BigEndian, &data.Size)
		binary.Read(buf, binary.BigEndian, &data.Wins)
		binary.Read(buf, binary.BigEndian, &data.Losses)

		unpacked.Teams = append(unpacked.Teams, data)
	}

	return
}
