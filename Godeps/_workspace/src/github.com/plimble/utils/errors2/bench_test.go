package errors2

import (
	"testing"
)

func BenchmarkNewError2(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewInternal("t", "test")
	}
}

func BenchmarkToError2(b *testing.B) {
	err := NewInternal("t", "test")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ToError2(err)
	}
}

func BenchmarkIsNotFound(b *testing.B) {
	err := NewNotFound("t", "test")
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		IsNotFound(err)
	}
}
