package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
)

const CallSignLen = 32
const MottoLen = 128
const ServerLen = 8
const MessageLen = 128
const HashLen = 64
const WorldSettingsSize = 30

const ReplayFileName = "20170701-1926-fun.rec"

type Duration struct {
	days  int
	hours int
	mins  int
	secs  int
	usecs int
}

type ReplayHeader struct {
	magic         uint32 // record file type identifier
	version       uint32 // record file version
	offset        uint32 // length of the full header
	filetime      int64  // amount of time in the file
	player        uint32 // player that saved this record file
	flagsSize     uint32 // size of the flags data
	worldSize     uint32 // size of world database
	callSign      []rune // player's callsign
	motto         []rune // player's motto
	ServerVersion []rune // BZFS protocol version
	appVersion    []rune // BZFS application version
	realHash      []rune // hash of worldDatabase
	worldSetting  []rune // the game settings

	// @TODO
	flags []rune // a list of the flags types
	world []rune // the world
}

func readRuneLength(buf *bytes.Buffer, length int) (arr []rune) {
	for i := 0; i < length; i++ {
		ltr, _ := binary.ReadUvarint(buf)

		// Once we've zero'd out, that means that there's no more data so skip
		// the remaining bits
		if ltr == 0 {
			buf.Next(length - i - 1)
			break
		}

		arr = append(arr, rune(ltr))
	}

	return
}

func loadHeader(buf *bytes.Buffer, header *ReplayHeader) {
	binary.Read(buf, binary.BigEndian, &header.magic)
	binary.Read(buf, binary.BigEndian, &header.version)
	binary.Read(buf, binary.BigEndian, &header.offset)
	binary.Read(buf, binary.BigEndian, &header.filetime)
	binary.Read(buf, binary.BigEndian, &header.player)
	binary.Read(buf, binary.BigEndian, &header.flagsSize)
	binary.Read(buf, binary.BigEndian, &header.worldSize)

	header.callSign = readRuneLength(buf, CallSignLen)
	header.motto = readRuneLength(buf, MottoLen)
	header.ServerVersion = readRuneLength(buf, ServerLen)
	header.appVersion = readRuneLength(buf, MessageLen)
	header.realHash = readRuneLength(buf, HashLen)
	header.worldSetting = readRuneLength(buf, 4+WorldSettingsSize)

	return
}

func loadTimestamp(timestamp int64) (time Duration) {
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
	var header ReplayHeader

	dat, _ := ioutil.ReadFile(ReplayFileName)
	buf := bytes.NewBuffer(dat)

	loadHeader(buf, &header)
	timeStruct := loadTimestamp(header.filetime)

	fmt.Printf("magic:     0x%04X\n", header.magic)
	fmt.Printf("replay:    version %d\n", header.version)
	fmt.Printf("offset:    %d\n", header.offset)
	fmt.Printf("length:    %-d days, %d hours, %d minutes, %d seconds, %d usecs\n", timeStruct.days, timeStruct.hours, timeStruct.mins, timeStruct.secs, timeStruct.usecs)

	// @TODO
	fmt.Printf("start:     ...\n")
	fmt.Printf("end:       ...\n")

	fmt.Printf("author:    %s  (%s)\n", string(header.callSign), string(header.motto))
	fmt.Printf("bzfs:      bzfs-%s\n", string(header.appVersion))
	fmt.Printf("protocol:  %.8s\n", string(header.ServerVersion))
	fmt.Printf("flagSize:  %d\n", header.flagsSize)
	fmt.Printf("worldSize: %d\n", header.worldSize)
	fmt.Printf("worldHash: %s\n", string(header.realHash))
}
