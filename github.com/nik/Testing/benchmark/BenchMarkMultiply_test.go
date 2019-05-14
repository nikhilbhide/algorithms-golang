package benchmark

import (
	"github.com/nik/Testing/numbers"
	"testing"
	"time"
)

type Test struct {
	num1   int
	num2   int
	num3   []int
	result int
}

func BenchmarkMultiply(b *testing.B) {
	testData := [] Test{Test{num1: 10, num2: 20, num3: [] int{20, 2, 3}, result: 1000}}
	for counter:=0; counter<b.N; counter++ {
		for _, value := range testData {
			numbers.MultiplyNumbers(append(value.num3, value.num1, value.num2)...)
		}
	}
}

func BenchmarkMultiply1(b *testing.B) {
	testData := [] Test{Test{num1: 10000, num2: 200000, num3: [] int{-99920, 20, 3000}, result: 1000}}
	for counter:=0; counter<b.N; counter++ {
		for _, value := range testData {
			numbers.MultiplyNumbers(append(value.num3, value.num1, value.num2)...)
			time.Sleep(100000)
		}
	}

	}

func BenchmarkMultiply2(b *testing.B) {
	testData := [] Test{Test{num1: 10000, num2: 200000, num3: [] int{-99920, 20, 3000}, result: 1000}}
	for counter:=0; counter<b.N; counter++ {
		for _, value := range testData {
			numbers.MultiplyNumbers(append(value.num3, value.num1, value.num2)...)

			time.Sleep(10000000)
		}
		time.Sleep(10000000000)
	}
}