package ch9

import (
	"hash/maphash"
	"testing"
)

func BenchmarkExample(b *testing.B) {
	//x := sha256.New()
	x := maphash.Hash{}
	testBytes := make([]byte, 1024)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		x.Reset()
		x.Write(testBytes)
		_ = x.Sum(nil)
	}
}
