package day06

import (
	"github.com/pedantic79/aoc2022go/framework"
)

const day uint = 6

func init() {
	if framework.CheckDayAndPart(day, 1) {
		framework.Results = append(framework.Results, RunPart1)
	}

	if framework.CheckDayAndPart(day, 2) {
		framework.Results = append(framework.Results, RunPart2)
	}
}

func RunPart1() framework.AoCResult {
	return framework.Timer(day, 1, parse, part1)
}

func RunPart2() framework.AoCResult {
	return framework.Timer(day, 2, parse, part2)
}

func unique(s string) bool {
	for i := range s {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				return false
			}
		}
	}

	return true
}

func parse(input string) string {
	return input
}

func solve(input string, width int) int {
	for i := 0; i < len(input)-width; i++ {
		if unique(input[i : i+width]) {
			return i + width
		}
	}

	return 0
}

func part1(input string) int {
	return solve(input, 4)
}

func part2(input string) int {
	return solve(input, 14)
}
