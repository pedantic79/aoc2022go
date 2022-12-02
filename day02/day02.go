package day02

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
)

const day uint = 2

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

type game struct {
	you, me int
}

func parse(input string) []game {
	nums := []game{}
	for _, line := range strings.Split(input, "\n") {
		nums = append(nums, game{int(line[0] - 'A' + 1), int(line[2] - 'X' + 1)})
	}

	return nums
}

func part1(games []game) int {
	sum := 0
	for _, g := range games {
		sum += g.me + 3*((4+g.me-g.you)%3)
	}

	return sum
}

func part2(games []game) int {
	sum := 0
	for _, g := range games {
		sum += ((g.you+g.me)%3 + 1) + 3*(g.me-1)
	}

	return sum
}
