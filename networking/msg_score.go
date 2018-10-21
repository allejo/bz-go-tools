package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgScorePacket struct {
	Type   string      `json:"type"`
	Scores []ScoreData `json:"scores"`
}

type ScoreData struct {
	PlayerID  uint8  `json:"playerID"`
	Wins      uint16 `json:"wins"`
	Losses    uint16 `json:"losses"`
	TeamKills uint16 `json:"teamKills"`
}

func (m *MsgScorePacket) Unpack(buf *bytes.Buffer) (packet MsgScorePacket) {
	packet.Type = "MsgScore"

	var i, count uint8
	binary.Read(buf, binary.BigEndian, &count)

	for i = 0; i < count; i++ {
		var score ScoreData

		binary.Read(buf, binary.BigEndian, &score.PlayerID)
		binary.Read(buf, binary.BigEndian, &score.Wins)
		binary.Read(buf, binary.BigEndian, &score.Losses)
		binary.Read(buf, binary.BigEndian, &score.TeamKills)

		packet.Scores = append(packet.Scores, score)
	}

	return
}
