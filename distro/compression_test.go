package distro

import (
	"bytes"
	"testing"
)

const (
	clearString = "This is the test string used for testing"
)

var gzipArray = []byte{31, 139, 8, 0, 0, 9, 110, 136, 0, 255, 10, 201, 200, 44, 86, 200, 44, 86, 40, 201, 72, 85, 40, 73, 45, 46, 81, 40, 46, 41, 202, 204, 75, 87, 40, 45, 78, 77, 81, 72, 203, 47, 2, 11, 102, 230, 165, 3, 2, 0, 0, 255, 255, 45, 117, 163, 29, 40, 0, 0, 0}

// Output from gzip
//0x1f, 0x8b, 0x08, 0x08, 0x0f, 0xb9, 0xed, 0x59, 0x00, 0x03, 0x6b, 0x6b, 0x00, 0x0b, 0xc9, 0xc8, 0x2c, 0x56, 0x00, 0xa2, 0x92, 0x8c, 0x54, 0x85, 0x92, 0xd4, 0xe2, 0x12, 0x85, 0xe2, 0x92, 0xa2, 0xcc, 0xbc, 0x74, 0x85, 0xd2, 0xe2, 0xd4, 0x14, 0x85, 0xb4, 0xfc, 0x22, 0xb0, 0x20, 0x50, 0x00, 0x00, 0x2d, 0x75, 0xa3, 0x1d, 0x28, 0x00, 0x00, 0x00}

//FIXME
var zlibArray = []byte{120, 156, 10, 201, 200, 44, 86, 200, 44, 86, 40, 201, 72, 85, 40, 73, 45, 46, 81, 40, 46, 41, 202, 204, 75, 87, 40, 45, 78, 77, 81, 72, 203, 47, 2, 11, 102, 230, 165, 3, 2, 0, 0, 255, 255, 42, 231, 14, 227}

func TestGzip(t *testing.T) {
	buf := bytes.Buffer{}
	buf.WriteString(clearString)

	comp := gzipTransformer{}

	enc := comp.Transform(buf)

	if bytes.Compare(enc.Bytes(), gzipArray) != 0 {
		t.Fatal("Encoded string ", enc.Bytes(), " is not ", gzipArray)
	}
}

func BenchmarkGzip(b *testing.B) {
	buf := bytes.Buffer{}
	buf.WriteString(clearString)

	comp := gzipTransformer{}

	for i := 0; i < 100; i++ {
		comp.Transform(buf)
	}
}

func TestZlib(t *testing.T) {
	buf := bytes.Buffer{}
	buf.WriteString(clearString)

	comp := zlibTransformer{}

	enc := comp.Transform(buf)

	if bytes.Compare(enc.Bytes(), zlibArray) != 0 {
		t.Fatal("Encoded string ", enc.Bytes(), " is not ", gzipArray)
	}
}

func BenchmarkZlib(b *testing.B) {
	buf := bytes.Buffer{}
	buf.WriteString(clearString)

	comp := zlibTransformer{}

	for i := 0; i < 100; i++ {
		comp.Transform(buf)
	}
}
