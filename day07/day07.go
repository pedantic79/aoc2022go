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

func parse(input string) map[string]int {
	dirs := map[string]int{}
	lines := strings.Split(input, "\n")
	cwd := []string{}

	for i := 0; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "$ cd") {
			line := strings.Split(lines[i], " ")

			if line[len(line)-1] == ".." {
				cwd = cwd[:len(cwd)-1]
			} else {
				cwd = append(cwd, line[len(line)-1])
			}
		} else if strings.HasPrefix(lines[i], "$ ls") {
			total := 0
			i++
			for i < len(lines) && !strings.HasPrefix(lines[i], "$") {
				if !strings.HasPrefix(lines[i], "dir ") {
					total += util.Atoi(strings.Split(lines[i], " ")[0])
				}
				i++
			}
			i--

			for end := len(cwd); end >= 1; end-- {
				dirs[strings.Join(cwd[:end], "/")] += total
			}
		}
	}

	// fmt.Println(dirs)
	return dirs
}

func part1(dirs map[string]int) int {
	total := 0
	for _, v := range dirs {
		if v < 100000 {
			total += v
		}
	}

	return total
}

func part2(dirs map[string]int) int {
	remaining := 30000000 - (70000000 - dirs["/"])

	min := math.MaxInt
	for _, v := range dirs {
		if v > remaining {
			min = util.Min(min, v)
		}
	}

	return min
}
