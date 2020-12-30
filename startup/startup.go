package startup

// Package startup only exists to separate initialization from tests

import (
	"flag"

	_ "github.com/pedantic79/aoc2022go/framework"
)

func init() {
	flag.Parse()
}
