package util

import (
	"crypto/rand"
	"encoding/hex"
)

func NewUUIDv4() string {
	var uuid [16]byte
	_, err := rand.Read(uuid[:])
  if err!=nil{
    panic(err)
  }
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	var buf [36]byte
	encodeHex(buf[:], uuid)

	return string(buf[:])
}

func encodeHex(dst []byte, uuid [16]byte) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}
