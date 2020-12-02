package intcode

import (
	"math/rand"
	"testing"
)

var header OperationHeader

func TestDecodeHeader(t *testing.T) {
	rawHeader := 12345
	header = DecodeHeader(rawHeader)

	if header.Operation != Operation(45) ||
		header.Op1Mode != AccessMode(3) ||
		header.Op2Mode != AccessMode(2) ||
		header.Op3Mode != AccessMode(1) {
		t.Fail()
	}

}

func BenchmarkDecodeHeader(b *testing.B) {
	for n := 0; n < b.N; n++ {
		header = DecodeHeader(rand.Intn(99999))
	}
}
