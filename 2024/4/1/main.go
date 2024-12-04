package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
)

func main() {
	gridStr := utils.ReadLines("../input.txt")

	grid := utils.ToGrid(gridStr)

	combinations := rotationCombinations()
	total := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			vector := utils.Vector{X: x, Y: y}
			cell := grid.GetCell(vector)
			if cell == 'X' {
				for _, direction := range combinations {
					match := true
					for i, char := range "XMAS" {
						vector := direction[i].Add(vector)
						if grid.IsInside(vector) {
							if grid.GetCell(vector) != char {
								match = false
								break
							}
						} else {
							match = false
							break
						}
					}

					if match {
						total += 1
					}
				}
			}
		}
	}
	fmt.Println(total)
}

func rotationCombinations() [][]utils.Vector {
	var all [][]utils.Vector

	all = append(all, straigtLineXMAS())

	var rotated []utils.Vector
	for _, a := range straigtLineXMAS() {
		rotated = append(rotated, a.Rotate90())
	}
	all = append(all, rotated)

	rotated = []utils.Vector{}
	for _, a := range straigtLineXMAS() {
		rotated = append(rotated, a.Rotate90().Rotate90())
	}
	all = append(all, rotated)

	rotated = []utils.Vector{}
	for _, a := range straigtLineXMAS() {
		rotated = append(rotated, a.Rotate90().Rotate90().Rotate90())
	}
	all = append(all, rotated)

	for _, r := range rotateMatrix() {
		var rotated []utils.Vector
		for _, a := range diagonalLineXMAS() {
			rotated = append(rotated, a.Mult(r))
		}
		all = append(all, rotated)
	}

	return all
}

func straigtLineXMAS() []utils.Vector {
	return []utils.Vector{{X: 0, Y: 0}, {X: 1, Y: 0}, {X: 2, Y: 0}, {X: 3, Y: 0}}
}

func diagonalLineXMAS() []utils.Vector {
	return []utils.Vector{{X: 0, Y: 0}, {X: 1, Y: 1}, {X: 2, Y: 2}, {X: 3, Y: 3}}
}

func rotateMatrix() []utils.Vector {
	return []utils.Vector{{X: 1, Y: 1}, {X: 1, Y: -1}, {X: -1, Y: 1}, {X: -1, Y: -1}}
}
