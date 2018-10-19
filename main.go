package main

import (
	"./networking"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

const ReplayFileName = "20170701-1926-fun.rec"

type Duration struct {
	days  int
	hours int
	mins  int
	secs  int
	usecs int
}

func calcDuration(timestamp int64) (time Duration) {
	secs := timestamp / 1000000

	time.days = int(secs / (24 * 60 * 60))
	secs %= 24 * 60 * 60

	time.hours = int(secs / (60 * 60))
	secs %= 60 * 60

	time.mins = int(secs / 60)
	secs %= 60

	time.secs = int(secs)
	time.usecs = int(timestamp % 1000000)

	return
}

func main() {
	dat, _ := ioutil.ReadFile(ReplayFileName)
	buf := bytes.NewBuffer(dat)

	header := networking.LoadReplayHeader(buf)
	timeStruct := calcDuration(header.FileTime)

	p, err := networking.LoadReplayPacket(buf)

	fmt.Printf("magic:     0x%04X\n", header.MagicNumber)
	fmt.Printf("replay:    version %d\n", header.Version)
	fmt.Printf("offset:    %d\n", header.Offset)
	fmt.Printf("length:    %-d days, %d hours, %d minutes, %d seconds, %d usecs\n", timeStruct.days, timeStruct.hours, timeStruct.mins, timeStruct.secs, timeStruct.usecs)
	fmt.Printf("start:     %s\n", time.Unix(p.Timestamp/1000000, 0))
	fmt.Printf("end:       %s\n", time.Unix((p.Timestamp+header.FileTime)/1000000, 0))
	fmt.Printf("author:    %s (%s)\n", header.CallSign, header.Motto)
	fmt.Printf("bzfs:      bzfs-%s\n", header.AppVersion)
	fmt.Printf("protocol:  %.8s\n", header.ServerVersion)
	fmt.Printf("flagSize:  %d\n", header.FlagsSize)
	fmt.Printf("worldSize: %d\n", header.WorldSize)
	fmt.Printf("worldHash: %s\n", header.RealHash)
	fmt.Printf("\n")

	for err == nil {
		var packet interface{}

		switch p.Mode {
		case networking.RealPacket:
			fmt.Printf("Real Packet\n")
			break

		case networking.StatePacket:
			fmt.Printf("State Packet\n")
			break

		case networking.UpdatePacket:
			fmt.Printf("Update Packet\n")
			break
		}

		packet = networking.UnpackNetworkPacket(p.Code, p.Data)

		marshaled, _ := json.Marshal(packet)
		fmt.Println(string(marshaled))

		// Move on to the next packet
		p, err = networking.LoadReplayPacket(buf)
	}
}
