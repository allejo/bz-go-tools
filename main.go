package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"time"
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
	callSign      []byte // player's callsign
	motto         []byte // player's motto
	ServerVersion []byte // BZFS protocol version
	appVersion    []byte // BZFS application version
	realHash      []byte // hash of worldDatabase
	worldSetting  []byte // the game settings

	// Information that is not being tracked right now

	//flags []byte // a list of the flags types
	//world []byte // the world
}

type ReplayPacket struct {
	next        *ReplayPacket
	prev        *ReplayHeader
	mode        uint16
	code        uint16
	len         uint32
	nextFilePos uint32
	prevFilePos uint32
	timestamp   int64
	data        []byte
}

func unpackString(buf *bytes.Buffer, length int) (dest []byte) {
	dest = make([]byte, length)
	io.ReadFull(buf, dest)
	dest = bytes.Trim(dest, "\x00")

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

	header.callSign = unpackString(buf, CallSignLen)
	header.motto = unpackString(buf, MottoLen)
	header.ServerVersion = unpackString(buf, ServerLen)
	header.appVersion = unpackString(buf, MessageLen)
	header.realHash = unpackString(buf, HashLen)
	header.worldSetting = unpackString(buf, 4+WorldSettingsSize)

	// Skip the appropriate number of bytes since we don't really care about this
	// data, for now

	if header.flagsSize > 0 {
		buf.Next(int(header.flagsSize))
	}

	buf.Next(int(header.worldSize))

	return
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

func loadPacket(buf *bytes.Buffer, packet *ReplayPacket) {
	binary.Read(buf, binary.BigEndian, &packet.mode)
	binary.Read(buf, binary.BigEndian, &packet.code)
	binary.Read(buf, binary.BigEndian, &packet.len)
	binary.Read(buf, binary.BigEndian, &packet.nextFilePos)
	binary.Read(buf, binary.BigEndian, &packet.prevFilePos)
	binary.Read(buf, binary.BigEndian, &packet.timestamp)

	if packet.len == 0 {
		packet.data = nil
	} else {
		packet.data = make([]byte, packet.len)
		io.ReadFull(buf, packet.data)
	}

	return
}

func main() {
	var header ReplayHeader

	dat, _ := ioutil.ReadFile(ReplayFileName)
	buf := bytes.NewBuffer(dat)

	loadHeader(buf, &header)
	timeStruct := calcDuration(header.filetime)

	var p ReplayPacket
	loadPacket(buf, &p)

	fmt.Printf("magic:     0x%04X\n", header.magic)
	fmt.Printf("replay:    version %d\n", header.version)
	fmt.Printf("offset:    %d\n", header.offset)
	fmt.Printf("length:    %-d days, %d hours, %d minutes, %d seconds, %d usecs\n", timeStruct.days, timeStruct.hours, timeStruct.mins, timeStruct.secs, timeStruct.usecs)
	fmt.Printf("start:     %s\n", time.Unix(p.timestamp/1000000, 0))
	fmt.Printf("end:       %s\n", time.Unix((p.timestamp+header.filetime)/1000000, 0))
	fmt.Printf("author:    %s (%s)\n", string(header.callSign), string(header.motto))
	fmt.Printf("bzfs:      bzfs-%s\n", string(header.appVersion))
	fmt.Printf("protocol:  %.8s\n", string(header.ServerVersion))
	fmt.Printf("flagSize:  %d\n", header.flagsSize)
	fmt.Printf("worldSize: %d\n", header.worldSize)
	fmt.Printf("worldHash: %s\n", string(header.realHash))
}
