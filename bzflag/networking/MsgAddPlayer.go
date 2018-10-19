package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgAddPlayerPacket struct {
	Type        string `json:"type"`
	PlayerIndex uint8  `json:"index"`
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

func (m *MsgAddPlayerPacket) Unpack(data []byte) (message MsgAddPlayerPacket) {
	buf := bytes.NewBuffer(data)

	message.Type = "MsgAddPlayer"

	binary.Read(buf, binary.BigEndian, &message.PlayerIndex)
	binary.Read(buf, binary.BigEndian, &message.PlayerType)
	binary.Read(buf, binary.BigEndian, &message.TeamValue)
	binary.Read(buf, binary.BigEndian, &message.Score.Wins)
	binary.Read(buf, binary.BigEndian, &message.Score.Losses)
	binary.Read(buf, binary.BigEndian, &message.Score.Teamkills)

	message.CallSign = UnpackString(buf, CallSignLen)
	message.Motto = UnpackString(buf, MottoLen)

	return
}
