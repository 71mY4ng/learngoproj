package learngo

import (
	"testing"
)

func BenchmarkStructReturnValue(b *testing.B) {

	b.ReportAllocs()

	t := 0
	for i := 0; i < b.N; i++ {
		v := newBigStruct()
		t += v.a[0]
	}
}

func BenchmarkStructReturnPtr(b *testing.B) {

	b.ReportAllocs()

	t := 0
	for i := 0; i < b.N; i++ {
		v := newBigStructPtr()
		t += v.a[0]
	}
}
