package day05

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
)

const day uint = 5

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

type moves struct {
	count, from_stack, to_stack int
}

func parseMoves(input string) []moves {
	m := []moves{}
	for _, line := range strings.Split(input, "\n") {
		fields := strings.Split(line, " ")

		m = append(m, moves{
			count:      util.Atoi(fields[1]),
			from_stack: util.Atoi(fields[3]) - 1,
			to_stack:   util.Atoi(fields[5]) - 1,
		})
	}

	return m
}

func parseCrates(input string) [][]byte {
	lines := strings.Split(input, "\n")
	size := (len(lines[len(lines)-1]) + 2) / 4
	stacks := make([][]byte, size)

	for x := range stacks {
		stacks[x] = make([]byte, 0, 64)
	}

	for i := len(lines) - 2; i >= 0; i-- {
		for j, stack_num := 0, 0; j < len(lines[i]); j += 4 {
			if lines[i][j+1] != ' ' {
				stacks[stack_num] = append(stacks[stack_num], lines[i][j+1])
			}
			stack_num++
		}
	}

	return stacks
}

type puzzle struct {
	stacks [][]byte
	m      []moves
}

func parse(input string) puzzle {
	chunks := strings.Split(input, "\n\n")
	return puzzle{
		parseCrates(chunks[0]),
		parseMoves(chunks[1]),
	}
}

func read_top(stacks [][]byte) string {
	s := strings.Builder{}

	for i := range stacks {
		s.WriteByte(stacks[i][len(stacks[i])-1])
	}

	return s.String()
}

func solve(input puzzle, reverse bool) string {
	for _, m := range input.m {
		var x []byte
		l := len(input.stacks[m.from_stack]) - m.count
		x, input.stacks[m.from_stack] = input.stacks[m.from_stack][l:], input.stacks[m.from_stack][:l]
		if reverse {
			util.Reverse(x)
		}
		input.stacks[m.to_stack] = append(input.stacks[m.to_stack], x...)
	}

	return read_top(input.stacks)
}

func part1(input puzzle) string {
	return solve(input, true)
}

func part2(input puzzle) string {
	return solve(input, false)
}
