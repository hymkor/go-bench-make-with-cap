package main

import (
	"fmt"
	"testing"
)

func BenchmarkAppendN(b *testing.B) {
	buffer := []string{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer = append(buffer, fmt.Sprintf("%d", i))
	}
}

func BenchmarkMakeWithCapAndAppend(b *testing.B) {
	buffer := make([]string, 0, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer = append(buffer, fmt.Sprintf("%d", i))
	}
}

func BenchmarkNoAppend(b *testing.B) {
	buffer := make([]string, b.N)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buffer[i] = fmt.Sprintf("%d", i)
	}
}
