package chunks

import (
	"math/rand"
	"reflect"
	"testing"
)

var s1 = randomString(500)
var s2 = randomString(4000)
var s3 = randomString(20000)

func randomString(n int) string {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func DoBenchmarks(f func(s string, chunkSize int) []string) {
	f(s1, 10)
	f(s2, 10)
	f(s3, 10)
	f(s1, 100)
	f(s2, 100)
	f(s3, 100)
	f(s1, 1000)
	f(s2, 1000)
	f(s3, 1000)
}

func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(Chunks)
	}
	b.ReportAllocs()
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(SplitSubN)
	}
	b.ReportAllocs()
}

func Benchmark3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(ChunkString)
	}
	b.ReportAllocs()
}

func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DoBenchmarks(ChunkStringImproved)
	}
	b.ReportAllocs()
}

func TestEquality(t *testing.T) {
	if !reflect.DeepEqual(SplitSubN(s3, 100), Chunks(s3, 100)) {
		t.Error()
	}
	if !reflect.DeepEqual(ChunkString(s3, 100), Chunks(s3, 100)) {
		t.Error()
	}
	if !reflect.DeepEqual(ChunkStringImproved(s3, 100), Chunks(s3, 100)) {
		t.Error()
	}
}
