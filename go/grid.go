package utils

type Grid[x comparable] struct {
	cells  []x
	Width  int
	Height int
}

func ToGrid(a []string) Grid[rune] {
	cells := make([]rune, 0)
	height := 0
	width := 0
	for y := 0; y < len(a); y++ {
		row := a[y]
		height = len(a)
		for x := 0; x < len(row); x++ {
			cells = append(cells, rune(row[x]))
			width = len(row)
		}
	}

	return Grid[rune]{
		cells:  cells,
		Width:  width,
		Height: height,
	}
}

func (g Grid[x]) GetCell(v Vector) x {
	return g.cells[g.ToIndex(v)]
}

func (g Grid[x]) ToIndex(v Vector) int {
	return g.Width*v.Y + v.X
}

func (g Grid[x]) IsInside(v Vector) bool {
	return v.X >= 0 && v.Y >= 0 && v.X < g.Width && v.Y < g.Height
}
