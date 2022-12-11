package day10

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
)

const day uint = 10

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

func parse(input string) []string {
	return strings.Split(input, "\n")
}

func parse_line(line string) (bool, int) {
	if line[0] == 'n' {
		return false, 0
	} else {
		return true, util.Atoi(line[5:])
	}
}

func solve(lines []string, update func(int, int)) {
	x := 1
	cycle := 1

	for _, line := range lines {
		update(cycle, x)
		if isAddx, count := parse_line(line); !isAddx {
			cycle++
		} else {
			cycle++
			update(cycle, x)
			cycle++
			x += count
		}
	}
}

func part1(lines []string) int {
	total := 0
	solve(lines, func(cycle, x int) {
		if cycle%40 == 20 {
			total += cycle * x
		}
	})

	return total
}

func draw(screen *strings.Builder, cycle, x int) {
	col := (cycle - 1) % 40

	if col == 0 {
		screen.WriteByte('\n')
	}

	if util.IntAbs(x-col) <= 1 {
		screen.WriteByte('#')
	} else {
		screen.WriteByte('.')
	}
}

func part2(lines []string) string {
	screen := strings.Builder{}
	solve(lines, func(cycle, x int) { draw(&screen, cycle, x) })

	return screen.String()
}
