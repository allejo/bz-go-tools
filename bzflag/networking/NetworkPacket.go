package networking

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

type Vector3F [3]float32

func UnpackFloat(buf *bytes.Buffer) float32 {
	var value uint32

	binary.Read(buf, binary.BigEndian, &value)

	return math.Float32frombits(value)
}

func UnpackString(buf *bytes.Buffer, length int) string {
	unpacked := make([]byte, length)
	io.ReadFull(buf, unpacked)

	// Remove any NULL characters
	unpacked = bytes.Trim(unpacked, "\x00")

	return string(unpacked)
}

func UnpackVector(buf *bytes.Buffer) (vector Vector3F) {
	vector[0] = UnpackFloat(buf)
	vector[1] = UnpackFloat(buf)
	vector[2] = UnpackFloat(buf)

	return
}
