package main

import (
	"./networking"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

const ReplayFileName = "20170701-1926-fun.rec"

type Replay struct {
	Header  networking.ReplayHeader `json:"header"`
	Packets []interface{}           `json:"packets"`
}

func main() {
	dat, _ := ioutil.ReadFile(ReplayFileName)
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
	fmt.Println(string(marshaled))
}
