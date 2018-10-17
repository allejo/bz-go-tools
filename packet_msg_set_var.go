package main

import (
	"bytes"
	"encoding/binary"
)

type SetVarData struct {
	count    uint16
	Type     string        `json:"type"`
	Settings []BZDBSetting `json:"settings"`
}

type BZDBSetting struct {
	Name    string `json:"name"`
	nameLen uint8

	Value    string `json:"value"`
	valueLen uint8
}

func handleMsgSetVar(data []byte) (unpacked SetVarData) {
	buf := bytes.NewBuffer(data)

	unpacked.Type = "MsgSetVar"

	binary.Read(buf, binary.BigEndian, &unpacked.count)

	var i uint16
	for i = 0; i < unpacked.count; i++ {
		var setting BZDBSetting

		binary.Read(buf, binary.BigEndian, &setting.nameLen)
		setting.Name = unpackString(buf, int(setting.nameLen))

		binary.Read(buf, binary.BigEndian, &setting.valueLen)
		setting.Value = unpackString(buf, int(setting.valueLen))

		unpacked.Settings = append(unpacked.Settings, setting)
	}

	return
}
