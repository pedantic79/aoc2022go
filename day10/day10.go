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

func solve[T any](lines []string, initial *T, update func(*T, int, int)) {
	x := 1
	cycle := 1

	for _, line := range lines {
		update(initial, cycle, x)
		if isAddx, count := parse_line(line); !isAddx {
			cycle++
		} else {
			cycle++
			update(initial, cycle, x)
			cycle++
			x += count
		}
	}
}

func part1(lines []string) int {
	total := 0
	solve(lines, &total, func(t *int, cycle, x int) {
		if cycle%40 == 20 {
			*t += cycle * x
		}
	})

	return total
}

func draw(screen *[][]byte, cycle, x int) {
	cycle -= 1
	row := cycle / 40
	col := cycle % 40

	if util.IntAbs(x-col) <= 1 {
		(*screen)[row][col] = '#'
	}
}

func part2(lines []string) string {
	screen := [][]byte{
		[]byte("........................................"),
		[]byte("........................................"),
		[]byte("........................................"),
		[]byte("........................................"),
		[]byte("........................................"),
		[]byte("........................................"),
	}

	solve(lines, &screen, draw)

	res := strings.Builder{}
	for i := range screen {
		res.WriteByte('\n')
		res.Write(screen[i])
	}

	return res.String()
}
