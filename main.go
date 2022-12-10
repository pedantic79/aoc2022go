package main

import (
	"fmt"

	"github.com/pedantic79/aoc2022go/framework"
	_ "github.com/pedantic79/aoc2022go/startup"

	_ "github.com/pedantic79/aoc2022go/day01"
	_ "github.com/pedantic79/aoc2022go/day02"
	_ "github.com/pedantic79/aoc2022go/day03"
	_ "github.com/pedantic79/aoc2022go/day04"
	_ "github.com/pedantic79/aoc2022go/day05"
	_ "github.com/pedantic79/aoc2022go/day06"
	_ "github.com/pedantic79/aoc2022go/day07"
	_ "github.com/pedantic79/aoc2022go/day08"
	_ "github.com/pedantic79/aoc2022go/day09"
)

func main() {
	fmt.Printf("🎄 Advent of Code 2022 - Day %02d\n\n", *framework.Day)

	for _, result := range framework.Results {
		fmt.Printf("%v", result())
	}
}
