package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
	"strconv"
)

func main() {
	input := utils.ReadLines("../input.txt")

	obstacles := make(map[string]bool)

	width := len(input[0])
	height := len(input)

	var pos utils.Vector
	var dir utils.Vector

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			char := input[y][x]
			if char == '#' {
				obstacles[posToKey(x, y, width)] = true
			}

			if char == '^' {
				pos = utils.Vector{X: x, Y: y}
				dir = utils.Vector{X: 0, Y: -1}
			}
		}
	}

	visited := make(map[string]bool)
	visited[vecToKey(pos, width)] = true

	for {
		newPos := pos.Add(dir)

		if newPos.X > width-1 || newPos.X < 0 || newPos.Y > height-1 || newPos.Y < 0 {
			break
		}

		if obstacles[vecToKey(newPos, width)] {
			dir = nextDir(dir)
			continue
		}

		pos = newPos
		visited[vecToKey(newPos, width)] = true
	}

	fmt.Println(len(visited))
}

func nextDir(dir utils.Vector) utils.Vector {
	if dir.X == 0 && dir.Y == -1 {
		return utils.Vector{X: 1, Y: 0}
	}

	if dir.X == 1 && dir.Y == 0 {
		return utils.Vector{X: 0, Y: 1}
	}

	if dir.X == 0 && dir.Y == 1 {
		return utils.Vector{X: -1, Y: 0}
	}

	if dir.X == -1 && dir.Y == 0 {
		return utils.Vector{X: 0, Y: -1}
	}

	panic(dir)
}

func vecToKey(vec utils.Vector, width int) string {
	return posToKey(vec.X, vec.Y, width)
}

func posToKey(x, y, width int) string {
	key := (y*width + x)
	return strconv.Itoa(key)
}
