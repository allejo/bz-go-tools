package main

import (
	"./networking"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

type Replay struct {
	Header  networking.ReplayHeader `json:"header"`
	Packets []interface{}           `json:"packets"`
}

func main() {
	replayFile := os.Args[1]
	dat, _ := ioutil.ReadFile(replayFile)
	buf := bytes.NewBuffer(dat)

	var replay Replay

	replay.Header = networking.LoadReplayHeader(buf)
	p, err := networking.LoadReplayPacket(buf)

	for err == nil {
		var packet interface{}

		packet = networking.UnpackNetworkPacket(p.Code, p.Data)
		replay.Packets = append(replay.Packets, packet)

		// Move on to the next packet
		p, err = networking.LoadReplayPacket(buf)
	}

	marshaled, _ := json.Marshal(replay)
	output := strings.Replace(replayFile, filepath.Ext(replayFile), ".json", 1)

	ioutil.WriteFile(output, marshaled, 0644)
}
