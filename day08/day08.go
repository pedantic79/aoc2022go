package day08

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
)

const day uint = 8

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

func parse(input string) [][]byte {
	grid := [][]byte{}
	for _, line := range strings.Split(input, "\n") {
		row := make([]byte, len(line))
		for i := 0; i < len(line); i++ {
			row[i] = line[i]
		}

		grid = append(grid, row)
	}

	return grid
}

func left(grid [][]byte, r, c int) []byte {
	ans := []byte{}
	for ci := c - 1; ci >= 0; ci-- {
		ans = append(ans, grid[r][ci])
	}
	return ans
}

func right(grid [][]byte, r, c int) []byte {
	return grid[r][(c + 1):]
}

func up(grid [][]byte, r, c int) []byte {
	ans := []byte{}
	for ri := r - 1; ri >= 0; ri-- {
		ans = append(ans, grid[ri][c])
	}
	return ans
}

func down(grid [][]byte, r, c int) []byte {
	ans := []byte{}
	for ri := r + 1; ri < len(grid); ri++ {
		ans = append(ans, grid[ri][c])
	}
	return ans
}

func all_less(slice []byte, max byte) bool {
	for i := range slice {
		if slice[i] >= max {
			return false
		}
	}

	return true
}

func check(grid [][]byte, r, c int) bool {
	height := grid[r][c]

	lx := all_less(left(grid, r, c), height)
	rx := all_less(right(grid, r, c), height)
	ux := all_less(up(grid, r, c), height)
	dx := all_less(down(grid, r, c), height)

	return lx || rx || ux || dx
}

func count_tree(slice []byte, height byte) int {
	count := 0
	for hi := range slice {
		count += 1
		if slice[hi] >= height {
			break
		}
	}

	return count
}

func calc_score(grid [][]byte, r, c int) int {
	height := grid[r][c]

	return count_tree(left(grid, r, c), height) * count_tree(right(grid, r, c), height) * count_tree(up(grid, r, c), height) * count_tree(down(grid, r, c), height)
}

func part1(grid [][]byte) int {
	count := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			if check(grid, r, c) {
				count++
			}
		}
	}

	return count
}

func part2(grid [][]byte) int {
	max := 0

	for r := 0; r < len(grid); r++ {
		for c := 0; c < len(grid[r]); c++ {
			max = util.Max(max, calc_score(grid, r, c))
		}
	}

	return max
}
