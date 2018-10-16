package main

import (
	"bytes"
	"encoding/binary"
)

type AddPlayerData struct {
	PlayerIndex uint8  `json:"index"`
	PlayerType  uint16 `json:"type"`
	TeamValue   uint16 `json:"team"`
	CallSign    string `json:"callsign"`
	Motto       string `json:"motto"`
	Score struct {
		Wins        uint16 `json:"wins"`
		Losses      uint16 `json:"losses"`
		Teamkills   uint16 `json:"teamkills"`
	} `json:"score"`
}

func handleMsgAddPlayer(len uint32, code uint16, data []byte) (unpacked AddPlayerData) {
	buf := bytes.NewBuffer(data)

	binary.Read(buf, binary.BigEndian, &unpacked.PlayerIndex)
	binary.Read(buf, binary.BigEndian, &unpacked.PlayerType)
	binary.Read(buf, binary.BigEndian, &unpacked.TeamValue)
	binary.Read(buf, binary.BigEndian, &unpacked.Score.Wins)
	binary.Read(buf, binary.BigEndian, &unpacked.Score.Losses)
	binary.Read(buf, binary.BigEndian, &unpacked.Score.Teamkills)

	unpacked.CallSign = unpackString(buf, CallSignLen)
	unpacked.Motto = unpackString(buf, MottoLen)

	return
}
