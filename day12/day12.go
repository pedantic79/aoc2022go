package day12

import (
	"strings"

	"github.com/pedantic79/aoc2022go/framework"
	"github.com/pedantic79/aoc2022go/util/set"
)

const day uint = 12

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

type coord struct {
	y, x int
}

type heightMap struct {
	m     [][]byte
	start coord
	end   coord
}

func parse(input string) heightMap {
	m := [][]byte{}
	start := coord{}
	end := coord{}

	lines := strings.Split(input, "\n")
	for r := range lines {
		l := []byte{}
		for c, x := range []byte(lines[r]) {
			if x == 'S' {
				start = coord{r, c}
				l = append(l, 'a')
			} else if x == 'E' {
				end = coord{r, c}
				l = append(l, 'z')
			} else {
				l = append(l, x)
			}
		}
		m = append(m, l)
	}

	return heightMap{m, start, end}
}

func bfs(start coord, successors func(coord) []coord, success func(coord) bool) int {
	if success(start) {
		return 1
	}
	depth := 0
	queue_now := []coord{}
	queue_new := []coord{}
	seen := set.New[coord]()

	queue_now = append(queue_now, start)
	set.Add(seen, start)

	for {
		depth++
		for j := range queue_now {
			next := successors(queue_now[j])
			for i := range next {
				if success(next[i]) {
					return depth
				}

				if !set.Contains(seen, next[i]) {
					queue_new = append(queue_new, next[i])
					set.Add(seen, next[i])
				}
			}
		}

		if len(queue_new) == 0 {
			break
		} else {
			queue_new, queue_now = queue_now[0:0], queue_new
		}
	}

	return -1
}

func part1(m heightMap) int {
	height := len(m.m)
	width := len(m.m[0])

	return bfs(m.start,
		func(c coord) []coord {
			neighbors := []coord{}
			h := m.m[c.y][c.x]
			if c.y != 0 && m.m[c.y-1][c.x] <= 1+h {
				neighbors = append(neighbors, coord{c.y - 1, c.x})
			}
			if c.y < height-1 && m.m[c.y+1][c.x] <= 1+h {
				neighbors = append(neighbors, coord{c.y + 1, c.x})
			}
			if c.x != 0 && m.m[c.y][c.x-1] <= 1+h {
				neighbors = append(neighbors, coord{c.y, c.x - 1})
			}
			if c.x < width-1 && m.m[c.y][c.x+1] <= 1+h {
				neighbors = append(neighbors, coord{c.y, c.x + 1})
			}

			return neighbors
		},
		func(c coord) bool {
			return c == m.end
		})
}

func part2(m heightMap) int {
	height := len(m.m)
	width := len(m.m[0])

	return bfs(m.end,
		func(c coord) []coord {
			neighbors := []coord{}
			h := m.m[c.y][c.x]
			if c.y != 0 && m.m[c.y-1][c.x]+1 >= h {
				neighbors = append(neighbors, coord{c.y - 1, c.x})
			}
			if c.y < height-1 && m.m[c.y+1][c.x]+1 >= h {
				neighbors = append(neighbors, coord{c.y + 1, c.x})
			}
			if c.x != 0 && m.m[c.y][c.x-1]+1 >= h {
				neighbors = append(neighbors, coord{c.y, c.x - 1})
			}
			if c.x < width-1 && m.m[c.y][c.x+1]+1 >= h {
				neighbors = append(neighbors, coord{c.y, c.x + 1})
			}

			return neighbors
		},
		func(c coord) bool {
			return m.m[c.y][c.x] == 'a'
		})
}
