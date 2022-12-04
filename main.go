package main

import (
	"fmt"

	_ "github.com/pedantic79/aoc2022go/day01"
	"github.com/pedantic79/aoc2022go/framework"
	_ "github.com/pedantic79/aoc2022go/startup"
)

func main() {
	fmt.Printf("🎄 Advent of Code 2022 - Day %02d\n\n", *framework.Day)

	for _, result := range framework.Results {
		fmt.Printf("%v", result())
	}
}
