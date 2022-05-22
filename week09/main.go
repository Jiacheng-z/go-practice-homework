package main

import (
	"encoding/binary"
	"fmt"
)

// 参考:
// https://github.com/Terry-Mao/goim/blob/master/examples/javascript/client.js

const (
	rawHeaderLen = 16
	packetOffset = 0
	headerOffset = 4
	verOffset    = 6
	opOffset     = 8
	seqOffset    = 12
)

func main() {
	data := encoder("Hello, World!")
	decoder(data)
}

func decoder(data []byte) {
	packetLen := binary.BigEndian.Uint32(data[packetOffset:headerOffset])
	headerLen := binary.BigEndian.Uint16(data[headerOffset:verOffset])
	version := binary.BigEndian.Uint16(data[verOffset:opOffset])
	operation := binary.BigEndian.Uint32(data[opOffset:seqOffset])
	sequence := binary.BigEndian.Uint32(data[seqOffset : seqOffset+4])

	fmt.Printf("packet: %v, header: %v, version: %v, op: %v, seq: %v\n", packetLen, headerLen, version, operation, sequence)
	body := string(data[headerLen:packetLen])
	fmt.Printf("body: %v\n", body)
}

func encoder(body string) []byte {
	headerLen := 16
	packetLen := len(body) + headerLen
	ret := make([]byte, packetLen)

	binary.BigEndian.PutUint32(ret[:4], uint32(packetLen))
	binary.BigEndian.PutUint16(ret[4:6], uint16(headerLen))

	version := 5
	binary.BigEndian.PutUint16(ret[6:8], uint16(version))
	operation := 6
	binary.BigEndian.PutUint32(ret[8:12], uint32(operation))
	sequence := 7
	binary.BigEndian.PutUint32(ret[12:16], uint32(sequence))

	byteBody := []byte(body)
	copy(ret[16:], byteBody)

	return ret
}
