//
// Copyright (c) 2017 Cavium
//
// SPDX-License-Identifier: Apache-2.0
//

package distro

import (
	"testing"
)

const (
	clearString = "This is the test string used for testing"
	gzipString  = "H4sIAAAJbogA/wrJyCxWyCxWKMlIVShJLS5RKC4pysxLVygtTk1RSMsvAgtm5qUDAgAA//8tdaMdKAAAAA=="
	zlibString  = "eJwKycgsVsgsVijJSFUoSS0uUSguKcrMS1coLU5NUUjLLwILZualAwIAAP//KucO4w=="
)

func TestGzip(t *testing.T) {

	comp := gzipTransformer{}
	enc := comp.Transform([]byte(clearString))

	if string(enc) != gzipString {
		t.Fatal("Encoded string ", string(enc), " is not ", gzipString)
	}
}

func TestZlib(t *testing.T) {

	comp := zlibTransformer{}
	enc := comp.Transform([]byte(clearString))
	if string(enc) != zlibString {
		t.Fatal("Encoded string ", string(enc), " is not ", zlibString)
	}
}

var result []byte

func BenchmarkGzip(b *testing.B) {

	comp := gzipTransformer{}

	var enc []byte
	for i := 0; i < 1000; i++ {
		enc = comp.Transform([]byte(clearString))
	}
	result = enc
}

func BenchmarkZlib(b *testing.B) {

	comp := zlibTransformer{}

	var enc []byte
	for i := 0; i < 1000; i++ {
		enc = comp.Transform([]byte(clearString))
	}
	result = enc
}
