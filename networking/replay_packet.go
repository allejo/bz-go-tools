package networking

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

type Duration struct {
	Days         int `json:"days"`
	Hours        int `json:"hours"`
	Minutes      int `json:"mins"`
	Seconds      int `json:"secs"`
	Milliseconds int `json:"usecs"`
}

type ReplayHeader struct {
	MagicNumber   uint32 `json:"magicNumber"`   // record file type identifier
	Version       uint32 `json:"version"`       // record file version
	Offset        uint32 `json:"offset"`        // length of the full header
	FileTime      int64  `json:"fileTime"`      // amount of time in the file
	Player        uint32 `json:"player"`        // player that saved this record file
	FlagsSize     uint32 `json:"flagSize"`      // size of the Flags data
	WorldSize     uint32 `json:"worldSize"`     // size of world database
	CallSign      string `json:"callsign"`      // player's callsign
	Motto         string `json:"motto"`         // player's motto
	ServerVersion string `json:"serverVersion"` // BZFS protocol version
	AppVersion    string `json:"appVersion"`    // BZFS application version
	RealHash      string `json:"realHash"`      // hash of worldDatabase
	WorldSetting  string `json:"worldSetting"`  // the game settings

	Length Duration `json:"length"`

	// Information that is not being tracked right now

	//Flags []byte // a list of the Flags types
	//world []byte // the world
}

type ReplayPacket struct {
	Mode        uint16
	Code        uint16
	Length      uint32
	NextFilePos uint32
	PrevFilePos uint32
	Timestamp   int64
	Data        []byte
}

func calcDuration(timestamp int64, length *Duration) {
	secs := timestamp / 1000000

	length.Days = int(secs / (24 * 60 * 60))
	secs %= 24 * 60 * 60

	length.Hours = int(secs / (60 * 60))
	secs %= 60 * 60

	length.Minutes = int(secs / 60)
	secs %= 60

	length.Seconds = int(secs)
	length.Milliseconds = int(timestamp % 1000000)

	return
}

func LoadReplayHeader(buf *bytes.Buffer) (header ReplayHeader) {
	binary.Read(buf, binary.BigEndian, &header.MagicNumber)
	binary.Read(buf, binary.BigEndian, &header.Version)
	binary.Read(buf, binary.BigEndian, &header.Offset)
	binary.Read(buf, binary.BigEndian, &header.FileTime)
	binary.Read(buf, binary.BigEndian, &header.Player)
	binary.Read(buf, binary.BigEndian, &header.FlagsSize)
	binary.Read(buf, binary.BigEndian, &header.WorldSize)

	header.CallSign = UnpackString(buf, CallSignLen)
	header.Motto = UnpackString(buf, MottoLen)
	header.ServerVersion = UnpackString(buf, ServerLen)
	header.AppVersion = UnpackString(buf, MessageLen)
	header.RealHash = UnpackString(buf, HashLen)
	header.WorldSetting = UnpackString(buf, 4+WorldSettingsSize)

	// Skip the appropriate number of bytes since we don't really care about this
	// data, for now

	if header.FlagsSize > 0 {
		buf.Next(int(header.FlagsSize))
	}

	buf.Next(int(header.WorldSize))

	calcDuration(header.FileTime, &header.Length)

	return
}

func LoadReplayPacket(fullBuffer *bytes.Buffer) (packet ReplayPacket, err error) {
	packetChunk := make([]byte, 32)
	_, err = io.ReadFull(fullBuffer, packetChunk)

	if err != nil {
		return
	}

	buf := bytes.NewBuffer(packetChunk)

	binary.Read(buf, binary.BigEndian, &packet.Mode)
	binary.Read(buf, binary.BigEndian, &packet.Code)
	binary.Read(buf, binary.BigEndian, &packet.Length)
	binary.Read(buf, binary.BigEndian, &packet.NextFilePos)
	binary.Read(buf, binary.BigEndian, &packet.PrevFilePos)
	binary.Read(buf, binary.BigEndian, &packet.Timestamp)

	if packet.Length == 0 {
		packet.Data = nil
	} else {
		packet.Data = make([]byte, packet.Length)
		binary.Read(fullBuffer, binary.BigEndian, &packet.Data)
	}

	return
}
