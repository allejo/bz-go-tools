package networking

import (
	"bytes"
	"encoding/binary"
)

const (
	IsRegistered = 1 << 0
	IsVerified   = 1 << 1
	IsAdmin      = 1 << 2
)

type MsgPlayerInfoPacket struct {
	Type    string       `json:"type"`
	Players []PlayerData `json:"players"`
}

type PlayerData struct {
	PlayerID     uint8 `json:"playerID"`
	IsRegistered bool  `json:"isRegistered"`
	IsVerified   bool  `json:"isVerified"`
	IsAdmin      bool  `json:"isAdmin"`
}

func (m *MsgPlayerInfoPacket) Unpack(buf *bytes.Buffer) (packet MsgPlayerInfoPacket) {
	packet.Type = "MsgPlayerInfo"

	var i, count uint8
	binary.Read(buf, binary.BigEndian, &count)

	for i = 0; i < count; i++ {
		var player PlayerData

		binary.Read(buf, binary.BigEndian, &player.PlayerID)

		var properties uint8
		binary.Read(buf, binary.BigEndian, &properties)

		player.IsRegistered = properties&IsRegistered == IsRegistered
		player.IsVerified = properties&IsVerified == IsVerified
		player.IsAdmin = properties&IsAdmin == IsAdmin

		packet.Players = append(packet.Players, player)
	}

	return
}
