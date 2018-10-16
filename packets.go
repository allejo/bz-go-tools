package main

import (
	"bytes"
	"encoding/binary"
	"io"
)

const CallSignLen = 32
const MottoLen = 128
const ServerLen = 8
const MessageLen = 128
const HashLen = 64
const WorldSettingsSize = 30

const (
	RealPacket   = 0 // broadcasted during replay
	StatePacket  = 1 // broadcasted to those you aren't yet stateful
	UpdatePacket = 2 // never broadcasted (only for replay use)
	HiddenPacket = 3 // never broadcasted (stored for admin. purposes)
)

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

	// Remove any NULL characters
	dest = bytes.Trim(dest, "\x00")

	return
}

func loadReplayHeader(buf *bytes.Buffer, header *ReplayHeader) {
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

func loadReplayPacket(buf *bytes.Buffer) (packet ReplayPacket, err error) {
	binary.Read(buf, binary.BigEndian, &packet.mode)
	binary.Read(buf, binary.BigEndian, &packet.code)
	binary.Read(buf, binary.BigEndian, &packet.len)
	binary.Read(buf, binary.BigEndian, &packet.nextFilePos)
	binary.Read(buf, binary.BigEndian, &packet.prevFilePos)
	binary.Read(buf, binary.BigEndian, &packet.timestamp)

	// The 2.4 protocol has an 8 byte padding for packets for some reason
	//   https://git.io/fxufC
	buf.Next(8)

	if packet.len == 0 {
		packet.data = nil
	} else {
		packet.data = make([]byte, packet.len)
		binary.Read(buf, binary.BigEndian, &packet.data)
	}

	return
}
