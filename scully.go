package main

import (
	"./networking"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type CaptureRecord struct {
	TeamCaptured  uint16
	TeamCapturing uint16
	Timestamp     int64
}

type KillRecord struct {
	Kills     uint
	Losses    uint
	TeamKills uint
}

type Player struct {
	CallSign string
	Motto    string
	Captures []CaptureRecord
	Enemies  map[string]*KillRecord

	isJoined bool
	cTeam    uint16
}

type Analysis struct {
	Players map[string]*Player

	roster map[uint8]*Player
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("please specify a replay file to parse...")
		panic(nil)
	}

	replayFile := os.Args[1]
	data, _ := ioutil.ReadFile(replayFile)
	buf := bytes.NewBuffer(data)

	networking.LoadReplayHeader(buf)
	p, err := networking.LoadReplayPacket(buf)

	var match Analysis
	match.Players = make(map[string]*Player)
	match.roster = make(map[uint8]*Player)

	for err == nil {
		var packet interface{}

		packet = networking.UnpackNetworkPacket(p.Code, p.Data)

		switch packet := packet.(type) {
		case networking.MsgAddPlayerPacket:
			var player *Player

			if _, exists := match.Players[packet.CallSign]; !exists {
				player = &Player{
					CallSign: packet.CallSign,
					Motto:    packet.Motto,
					Captures: []CaptureRecord{},
					Enemies:  make(map[string]*KillRecord),
				}

				match.Players[packet.CallSign] = player
			} else {
				player = match.Players[packet.CallSign]
			}

			player.isJoined = true
			player.cTeam = packet.TeamValue

			match.roster[packet.PlayerIndex] = player

		case networking.MsgPlayerInfoPacket:
			for _, player := range packet.Players {
				match.roster[player.PlayerID].isJoined = player.IsVerified
			}

		case networking.MsgRemovePlayerPacket:
			match.roster[packet.PlayerID].isJoined = false

		case networking.MsgKilledPacket:
			victim := match.roster[packet.VictimID]
			killer := match.roster[packet.KillerID]

			if _, exists := victim.Enemies[killer.CallSign]; !exists {
				victim.Enemies[killer.CallSign] = &KillRecord{}
			}

			if _, exists := killer.Enemies[victim.CallSign]; !exists {
				killer.Enemies[victim.CallSign] = &KillRecord{}
			}

			victim.Enemies[killer.CallSign].Losses++
			killer.Enemies[victim.CallSign].Kills++

		case networking.MsgCaptureFlagPacket:
			capper := match.roster[packet.PlayerID]
			record := CaptureRecord{
				TeamCapturing: capper.cTeam,
				Timestamp:     p.Timestamp,
			}
			capper.Captures = append(capper.Captures, record)
			break

		default:
			break
		}

		p, err = networking.LoadReplayPacket(buf)
	}

	m, err := json.Marshal(match)
	fmt.Println(string(m))
}
