package networking

import (
	"bytes"
	"encoding/binary"
)

type MsgSetVarPacket struct {
	count    uint16
	Type     string        `json:"type"`
	Settings []BZDBSetting `json:"settings"`
}

type BZDBSetting struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (m *MsgSetVarPacket) Unpack(buf *bytes.Buffer) (packet MsgSetVarPacket) {
	packet.Type = "MsgSetVar"

	binary.Read(buf, binary.BigEndian, &packet.count)

	var i uint16
	for i = 0; i < packet.count; i++ {
		var setting BZDBSetting

		var nameLen uint8
		binary.Read(buf, binary.BigEndian, &nameLen)
		setting.Name = UnpackString(buf, int(nameLen))

		var valueLen uint8
		binary.Read(buf, binary.BigEndian, &valueLen)
		setting.Value = UnpackString(buf, int(valueLen))

		packet.Settings = append(packet.Settings, setting)
	}

	return
}
