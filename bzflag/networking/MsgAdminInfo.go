package networking

import (
	"bytes"
	"encoding/binary"
	"net"
)

type MsgAdminInfoPacket struct {
	Type string `json:"type"`
	Players []PlayerInfo `json:"players"`
}

type PlayerInfo struct {
	PlayerIndex uint8  `json:"playerID"`
	IPAddress   net.IP `json:"ipAddress"`
}

func (m *MsgAdminInfoPacket) Unpack(buf *bytes.Buffer) (packet MsgAdminInfoPacket) {
	var count, i, ipSize uint8

	packet.Type = "MsgAdminInfo"

	binary.Read(buf, binary.BigEndian, &count)

	for i = 0; i < count; i++ {
		binary.Read(buf, binary.BigEndian, &ipSize)

		var playerData PlayerInfo
		binary.Read(buf, binary.BigEndian, &playerData.PlayerIndex)

		playerData.IPAddress = UnpackIpAddress(buf)

		packet.Players = append(packet.Players, playerData)
	}

	return
}
