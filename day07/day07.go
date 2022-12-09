package day07

import (
	"math"
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
)

const day uint = 7

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

func calculate_sizes(lines []string, i *int, sizes *[]int) int {
	total := 0

	for *i < len(lines) {
		line := lines[*i]
		*i++
		if strings.HasPrefix(line, "$ cd") {
			line := strings.Split(line, " ")
			dir := line[len(line)-1]

			if dir == ".." {
				break
			} else if dir != "/" {
				total += calculate_sizes(lines, i, sizes)
			}
		} else if strings.HasPrefix(line, "$ ls") {
			for ; *i < len(lines) && !strings.HasPrefix(lines[*i], "$"); *i++ {
				if !strings.HasPrefix(lines[*i], "dir ") {
					total += util.Atoi(strings.Split(lines[*i], " ")[0])
				}
			}
			*i--
		}
	}

	*sizes = append(*sizes, total)
	return total
}

func parse(input string) []int {
	lines := strings.Split(input, "\n")
	res := []int{}
	i := 0
	calculate_sizes(lines, &i, &res)
	return res
}

func part1(dirs []int) int {
	total := 0
	for _, v := range dirs {
		if v < 100000 {
			total += v
		}
	}

	return total
}

func part2(dirs []int) int {
	remaining := 30000000 - (70000000 - dirs[len(dirs)-1])

	min := math.MaxInt
	for _, v := range dirs {
		if v > remaining {
			min = util.Min(min, v)
		}
	}

	return min
}
