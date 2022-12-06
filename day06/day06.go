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

func unique(s string) (int, bool) {
	for i := len(s) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				return j + 1, false
			}
		}
	}

	return 0, true
}

func parse(input string) string {
	return input
}

func solve(input string, width int) int {
	for i := 0; i < len(input)-width; {
		if j, flag := unique(input[i : i+width]); !flag {
			i += j
		} else {
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
