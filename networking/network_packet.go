package networking

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"net"
)

type Vector3F [3]float32

func UnpackFloat(buf *bytes.Buffer) float32 {
	var value uint32

	binary.Read(buf, binary.BigEndian, &value)

	return math.Float32frombits(value)
}

func UnpackIpAddress(buf *bytes.Buffer) net.IP {
	// This byte was reserved for differentiating between IPv4 and IPv6 addresses
	// However, since BZFlag only supports IPv4, this byte is skipped
	buf.Next(1)

	// IP Addresses are stored in network byte order (aka Little Endian)
	var ipAddress uint32
	binary.Read(buf, binary.LittleEndian, &ipAddress)

	ip := make(net.IP, 4)
	binary.LittleEndian.PutUint32(ip, ipAddress)

	return ip
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
