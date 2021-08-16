package strByte

import (
	"unsafe"
)

func String2Bytes(str string) []byte {
	return *(*[]byte)(unsafe.Pointer(&struct{ string }{str}))
}

func Bytes2String(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
