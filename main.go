package main

import (
	"fmt"

	lib "github.com/mupersega/arrayPlay/library"
)

func main() {
	m := lib.CreateMatrix(20, 20)

	lib.Display2DArray(m.SeedData)
	fmt.Println()
	lib.Display2DArray(m.Smoothed)
}
