package day13

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
)

const day uint = 13

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

func parse(input string) []any {
	nums := []any{}

	for _, line := range strings.Split(input, "\n") {
		if len(line) > 0 {
			var v any
			json.Unmarshal([]byte(line), &v)
			nums = append(nums, v)
		}
	}

	return nums
}

func getJsonNumOrList(v any) ([]any, bool) {
	switch e := v.(type) {
	case float64:
		return []any{e}, true
	case []any:
		return e, false
	default:
		panic(fmt.Sprintf("unknown type: %t", v))
	}

}

func cmpJsonValue(a, b any) int {
	sliceA, isFloatA := getJsonNumOrList(a)
	sliceB, isFloatB := getJsonNumOrList(b)

	if isFloatA && isFloatB {
		return int(sliceA[0].(float64) - sliceB[0].(float64))
	}

	return cmpJsonArray(sliceA, sliceB)
}

func cmpJsonArray(a, b []any) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if c := cmpJsonValue(a[i], b[i]); c != 0 {
			return c
		}
	}
	return len(a) - len(b)
}

func part1(signals []any) int {
	value := 0
	for i := 0; i < len(signals); i += 2 {
		if cmpJsonValue(signals[i], signals[i+1]) < 0 {
			value += i/2 + 1
		}
	}

	return value
}

func part2(signals []any) int {
	var two any
	var six any
	json.Unmarshal([]byte("[[2]]"), &two)
	json.Unmarshal([]byte("[[6]]"), &six)

	twoCount := 1
	sixCount := 2
	for i := range signals {
		if cmpJsonValue(signals[i], six) < 0 {
			sixCount++
			if cmpJsonValue(signals[i], two) < 0 {
				twoCount++
			}
		}
	}

	return sixCount * twoCount
}
