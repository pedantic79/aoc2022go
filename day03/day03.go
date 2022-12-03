package day03

import (
	"math/bits"
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
)

const day uint = 3

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
	lines := []string{}
	for _, line := range strings.Split(input, "\n") {
		lines = append(lines, line)
	}

	return lines
}

func priority(b rune) uint64 {
	if 'a' <= b && b <= 'z' {
		return uint64(b-'a') + 1
	} else {
		return uint64(b-'A'+26) + 1
	}
}

func buildSet(s string) (set uint64) {
	for _, c := range s {
		set |= 1 << priority(c)
	}
	return
}

func part1(lines []string) uint64 {
	sum := uint64(0)
	for _, line := range lines {
		mid := len(line) / 2

		set := buildSet(line[:mid])
		for _, c := range line[mid:] {
			prio := priority(c)
			if set&(1<<prio) > 0 {
				sum += prio
				break
			}
		}
	}

	return sum
}

func part2(nums []string) uint64 {
	var sum uint64
	for i := 0; i < len(nums); i += 3 {
		sum += uint64(bits.TrailingZeros64(buildSet(nums[i]) & buildSet(nums[i+1]) & buildSet(nums[i+2])))
	}
	return sum
}
