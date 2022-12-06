package day02

import (
	"testing"

	"github.com/pedantic79/aoc2022go/framework"
)

func TestRegression(t *testing.T) {
	right := []int{13682, 12881}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}

func BenchmarkParse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		framework.BenchInput(day, parse)
	}
}

func BenchmarkPart1(b *testing.B) {
	input := framework.BenchInput(day, parse)
	for i := 0; i < b.N; i++ {
		framework.Bench(input, part1)
	}
}

func BenchmarkPart2(b *testing.B) {
	input := framework.BenchInput(day, parse)
	for i := 0; i < b.N; i++ {
		framework.Bench(input, part2)
	}
}
