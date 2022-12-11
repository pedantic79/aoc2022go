package day11

import (
	"sort"
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
)

const day uint = 11

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

type Monkey struct {
	items      []int
	op         func(int) int
	divisor    int
	test_true  int
	test_false int
}

func parseTrailingNum(line string) int {
	i := strings.LastIndex(line, " ") + 1
	return util.Atoi(line[i:])
}

func parseItems(line string) []int {
	i := strings.Index(line, ": ") + 1
	nums := strings.Split(line[i:], ", ")

	res := []int{}
	for i := range nums {
		res = append(res, util.Atoi(strings.TrimSpace(nums[i])))
	}

	return res
}

func parseOp(line string) func(int) int {
	i := strings.Index(line, " = ") + 1
	symbols := strings.Split(line[i:], " ")
	if symbols[3] == "old" {
		return func(x int) int { return x * x }
	}
	num := util.Atoi(symbols[3])
	if symbols[2] == "*" {
		return func(x int) int { return num * x }
	}

	return func(x int) int { return num + x }
}

func parseMonkey(chunk string) Monkey {
	lines := strings.Split(chunk, "\n")

	return Monkey{
		parseItems(lines[1]),
		parseOp(lines[2]),
		parseTrailingNum(lines[3]),
		parseTrailingNum(lines[4]),
		parseTrailingNum(lines[5]),
	}
}

func parse(input string) []Monkey {
	monkeys := []Monkey{}
	for _, chunk := range strings.Split(input, "\n\n") {
		monkeys = append(monkeys, parseMonkey(chunk))
	}

	return monkeys
}

func solve(monkeys []Monkey, limit int, mainterer func(int) int) int {
	inspects := make([]int, len(monkeys))
	for iteration := 0; iteration < limit; iteration++ {
		for i := range monkeys {
			for j := range monkeys[i].items {
				item := monkeys[i].items[j]
				worry := mainterer(monkeys[i].op(item))
				if worry%monkeys[i].divisor == 0 {
					monkeys[monkeys[i].test_true].items = append(monkeys[monkeys[i].test_true].items, worry)
				} else {
					monkeys[monkeys[i].test_false].items = append(monkeys[monkeys[i].test_false].items, worry)
				}
			}
			inspects[i] += len(monkeys[i].items)
			monkeys[i].items = monkeys[i].items[0:0]
		}
	}

	sort.Ints(inspects)
	return inspects[len(inspects)-1] * inspects[len(inspects)-2]
}

func part1(monkeys []Monkey) int {
	return solve(monkeys, 20, func(x int) int { return x / 3 })
}

func part2(monkeys []Monkey) int {
	product := 1
	for i := range monkeys {
		product *= monkeys[i].divisor
	}

	return solve(monkeys, 10000, func(x int) int { return x % product })
}
