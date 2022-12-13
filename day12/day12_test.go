package day12

import "testing"

func TestRegression(t *testing.T) {
	right := []int{423, 416}
	left := []interface{}{RunPart1().Value, RunPart2().Value}

	for i, l := range left {
		if right[i] != l {
			t.Errorf("Expected (%v), Actual (%v)", right[i], l)
		}
	}

}

func BenchmarkPart1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunPart1()
	}
}

func BenchmarkPart2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RunPart2()
	}
}
