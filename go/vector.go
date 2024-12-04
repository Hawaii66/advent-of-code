package utils

import "fmt"

type Vector struct {
	X int
	Y int
}

func (a Vector) String() string {
	return fmt.Sprintf("V(X: %d, Y: %d)", a.X, a.Y)
}

func (a Vector) Add(b Vector) Vector {
	return Vector{
		X: a.X + b.X,
		Y: a.Y + b.Y,
	}
}

func (a Vector) Sub(b Vector) Vector {
	return Vector{
		X: a.X - b.X,
		Y: a.Y - b.Y,
	}
}

func (a Vector) Mult(b Vector) Vector {
	return Vector{
		X: a.X * b.X,
		Y: a.Y * b.Y,
	}
}

func (a Vector) Flip() Vector {
	return Vector{
		X: a.Y,
		Y: a.X,
	}
}

func (a Vector) Rotate90() Vector {
	if a.X == 0 && a.Y == 0 {
		return Vector{X: 0, Y: 0}
	}

	if a.X >= 1 && a.Y == 0 {
		return Vector{X: 0, Y: -a.X}
	}
	if a.X == 0 && a.Y <= -1 {
		return Vector{X: a.Y, Y: 0}
	}
	if a.X <= -1 && a.Y == 0 {
		return Vector{X: 0, Y: -a.X}
	}
	if a.X == 0 && a.Y >= 1 {
		return Vector{X: a.Y, Y: 0}
	}

	fmt.Println("Tried to rotate vector", a)
	panic(a)
}
