package day09

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
	"github.com/pedantic79/aoc2022go/util/set"
)

const day uint = 9

func init() {
	if framework.CheckDayAndPart(day, 1) {
		framework.Results = append(framework.Results, RunPart1)
	}

	if framework.CheckDayAndPart(day, 2) {
		framework.Results = append(framework.Results, RunPart2)
	}
}

type move struct {
	dir    byte
	amount int
}

func RunPart1() framework.AoCResult {
	return framework.Timer(day, 1, parse, part1)
}

func RunPart2() framework.AoCResult {
	return framework.Timer(day, 2, parse, part2)
}

func parse(input string) []move {
	moves := []move{}
	lines := strings.Split(input, "\n")

	for _, line := range lines {
		dir := line[0]
		amount := util.Atoi(line[2:])

		moves = append(moves, move{dir, amount})
	}

	return moves
}

type pos struct {
	y, x int
}

func step(prev pos, current *pos) {
	diff := pos{prev.y - current.y, prev.x - current.x}

	if util.IntAbs(diff.y) > 1 || util.IntAbs(diff.x) > 1 {
		current.y += util.SigNum(diff.y)
		current.x += util.SigNum(diff.x)
	}
}

var delta_row = map[byte]int{'U': 1, 'D': -1, 'L': 0, 'R': 0}
var delta_col = map[byte]int{'U': 0, 'D': 0, 'L': -1, 'R': 1}

func solve(moves []move, size int) int {
	rope := make([]pos, size)
	seen := set.New[pos]()
	for i := range moves {
		for j := 0; j < moves[i].amount; j++ {
			rope[0].y += delta_row[moves[i].dir]
			rope[0].x += delta_col[moves[i].dir]

			for tail := 1; tail < size; tail++ {
				step(rope[tail-1], &rope[tail])
			}
			set.Add(seen, rope[size-1])
		}
	}

	return len(seen)
}

func part1(moves []move) int {
	return solve(moves, 2)
}

func part2(moves []move) int {
	return solve(moves, 10)
}
