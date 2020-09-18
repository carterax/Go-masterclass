package main

import (
	"strconv"
	"testing"
)

const x = int16(123)

func Benchmark_Itoa(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.Itoa(int(x))
	}
}

func Benchmark_FormatInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strconv.FormatInt(int64(x), 10)
	}
}
