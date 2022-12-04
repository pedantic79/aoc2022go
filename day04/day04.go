package day04

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
)

const day uint = 4

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

type assignments struct {
	a, b, x, y int
}

func parse(input string) []assignments {
	assigns := []assignments{}

	for _, line := range strings.Split(input, "\n") {
		overlap := strings.FieldsFunc(line, func(r rune) bool { return r == ',' || r == '-' })

		assigns = append(assigns, assignments{
			util.Atoi(overlap[0]),
			util.Atoi(overlap[1]),
			util.Atoi(overlap[2]),
			util.Atoi(overlap[3]),
		})
	}

	return assigns
}

func overlap1(a assignments) bool {
	return (a.a <= a.x && a.b >= a.y) || (a.x <= a.a && a.y >= a.b)
}

func overlap2(a assignments) bool {
	return a.a <= a.y && a.b >= a.x
}

func solve(assigns []assignments, pred func(assignments) bool) int {
	count := 0

	for _, a := range assigns {
		if pred(a) {
			count++
		}
	}

	return count
}

func part1(assigns []assignments) int {
	return solve(assigns, overlap1)
}

func part2(assigns []assignments) int {
	return solve(assigns, overlap2)
}
