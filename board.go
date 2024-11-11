package main

type Board struct {
	width, height       int
	currentCatsPosition []CatBackend
}

func NewBoard(width, height int) *Board {
	cat1 := CatBackend{color: Green, X: 10, Y: 10}
	cat2 := CatBackend{color: Red, X: 20, Y: 15}

	return &Board{width, height, []CatBackend{cat1, cat2}}
}
