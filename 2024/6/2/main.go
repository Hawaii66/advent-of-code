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
	var startingPos utils.Vector
	var startingDir utils.Vector

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			char := input[y][x]
			if char == '#' {
				obstacles[posToKey(x, y, width)] = true
			}

			if char == '^' {
				pos = utils.Vector{X: x, Y: y}
				startingPos = utils.Vector{X: x, Y: y}
				dir = utils.Vector{X: 0, Y: -1}
				startingDir = utils.Vector{X: 0, Y: -1}
			}
		}
	}

	visited := make(map[string]utils.Vector)

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
		visited[vecToKey(newPos, width)] = utils.Vector{X: newPos.X, Y: newPos.Y}
	}

	total := 0
	for _, visitedPos := range visited {
		temp := make(map[string]bool)

		for key, value := range obstacles {
			temp[key] = value
		}

		temp[vecToKey(visitedPos, width)] = true
		if isLoop(startingPos, startingDir, width, height, temp) {
			total += 1
		}
	}

	fmt.Println(total)
}

func isLoop(pos, dir utils.Vector, width, height int, obstacles map[string]bool) bool {
	visited := make(map[string]bool)

	for {
		newPos := pos.Add(dir)
		key := posAndDirToKey(newPos, dir, width)
		if visited[key] {
			return true
		}

		if newPos.X > width-1 || newPos.X < 0 || newPos.Y > height-1 || newPos.Y < 0 {
			return false
		}

		if obstacles[vecToKey(newPos, width)] {
			dir = nextDir(dir)
			continue
		}

		pos = newPos
		visited[key] = true
	}
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

func posAndDirToKey(pos, dir utils.Vector, width int) string {
	return vecToKey(pos, width) + " - " + dir.String()
}

func vecToKey(vec utils.Vector, width int) string {
	return posToKey(vec.X, vec.Y, width)
}

func posToKey(x, y, width int) string {
	key := (y*width + x)
	return strconv.Itoa(key)
}
