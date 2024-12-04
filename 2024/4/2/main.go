package main

import (
	"fmt"
	utils "hawaiidev/advent-of-code/go"
	"strings"
)

func main() {
	gridStr := utils.ReadLines("../input.txt")

	grid := utils.ToGrid(gridStr)

	total := 0
	for x := 0; x < grid.Width; x++ {
		for y := 0; y < grid.Height; y++ {
			vector := utils.Vector{X: x, Y: y}
			cell := grid.GetCell(vector)
			match := true
			if cell == 'A' {
				var offsets []utils.Vector
				for _, o := range diagonalOffsets() {
					v := o.Add(vector)
					if !grid.IsInside(v) {
						match = false
						break
					}
					offsets = append(offsets, v)
				}

				if !match {
					continue
				}

				chars := ""
				for _, o := range offsets {
					char := grid.GetCell(o)
					chars += string(char)
				}

				if strings.Contains("MMSS", chars) || strings.Contains("MSSM", chars) || strings.Contains("SMMS", chars) || strings.Contains("SSMM", chars) {
					total += 1
				}
			}
		}
	}

	fmt.Println(total)
}

func diagonalOffsets() []utils.Vector {
	return []utils.Vector{
		{X: -1, Y: 1},
		{X: 1, Y: 1},
		{X: 1, Y: -1},
		{X: -1, Y: -1},
	}
}
