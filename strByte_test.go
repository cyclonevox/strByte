package strByte

import (
	"bytes"
	"testing"
)

func TestString2Bytes(t *testing.T) {
	test := "test,string"
	got := String2Bytes(test)
	want := []byte(test)

	if !bytes.Equal(got, want) {
		t.Errorf("got '%q' want '%q'", got, want)
	}

}

func TestBytes2String2(t *testing.T) {
	test := []byte("test,string")
	got := Bytes2String(test)
	want := "test,string"

	if want != got {
		t.Errorf("got '%q' want '%q'", got, want)
	}
}

func convert2BytesByOrigin(bytes []byte) string {
	return string(bytes)
}

func convert2StringByOrigin(str string) []byte {
	return []byte(str)
}

func BenchmarkConvert2BytesByOrigin(b *testing.B) {
	test := []byte("test,string,test,string,test,string,test,string")
	for i := 0; i < b.N; i++ {
		convert2BytesByOrigin(test)
	}
}

func BenchmarkConvert2StringByOrigin(b *testing.B) {
	test := "test,string,test,string,test,string,test,string"
	for i := 0; i < b.N; i++ {
		convert2StringByOrigin(test)
	}
}

func BenchmarkString2Bytes(b *testing.B) {
	test := "test,string,test,string,test,string,test,string"
	for i := 0; i < b.N; i++ {
		String2Bytes(test)
	}
}

func BenchmarkBytes2String(b *testing.B) {
	test := []byte("test,string,test,string,test,string,test,string")
	for i := 0; i < b.N; i++ {
		Bytes2String(test)
	}
}
