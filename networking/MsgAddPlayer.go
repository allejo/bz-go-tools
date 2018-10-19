package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgAddPlayerPacket struct {
	Type        string `json:"type"`
	PlayerIndex uint8  `json:"playerID"`
	PlayerType  uint16 `json:"type"`
	TeamValue   uint16 `json:"team"`
	CallSign    string `json:"callsign"`
	Motto       string `json:"motto"`
	Score       struct {
		Wins      uint16 `json:"wins"`
		Losses    uint16 `json:"losses"`
		Teamkills uint16 `json:"teamkills"`
	} `json:"score"`
}

func (m *MsgAddPlayerPacket) Unpack(buf *bytes.Buffer) (packet MsgAddPlayerPacket) {
	packet.Type = "MsgAddPlayer"

	binary.Read(buf, binary.BigEndian, &packet.PlayerIndex)
	binary.Read(buf, binary.BigEndian, &packet.PlayerType)
	binary.Read(buf, binary.BigEndian, &packet.TeamValue)
	binary.Read(buf, binary.BigEndian, &packet.Score.Wins)
	binary.Read(buf, binary.BigEndian, &packet.Score.Losses)
	binary.Read(buf, binary.BigEndian, &packet.Score.Teamkills)

	packet.CallSign = UnpackString(buf, CallSignLen)
	packet.Motto = UnpackString(buf, MottoLen)

	return
}
