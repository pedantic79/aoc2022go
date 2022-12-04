package day01

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util"
)

const day uint = 1

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

func parse(input string) []int {
	nums := []int{}

	for _, elf := range strings.Split(input, "\n\n") {
		sum := 0
		for _, amount := range strings.Split(elf, "\n") {
			sum += util.Atoi(amount)
		}
		nums = append(nums, sum)
	}

	return nums
}

func part1(nums []int) int {
	max := 0

	for _, num := range nums {
		max = util.Max(num, max)
	}
	return max
}

func part2(nums []int) int {
	const LEN int = 3
	var maxes [LEN]int

	for _, num := range nums {
		if num > maxes[LEN-1] {
			for i := 0; i < LEN; i++ {
				if num > maxes[i] {
					for j := LEN - 1; j > i; j-- {
						maxes[j] = maxes[j-1]
					}
					maxes[i] = num
					break
				}
			}
		}
	}

	return maxes[0] + maxes[1] + maxes[2]
}
