package library

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

type Shape struct {
	height int
	width  int
}

type Matrix struct {
	Shape    Shape
	SeedData [][]int
	Smoothed [][]int
}

func CreateMatrix(height int, width int) Matrix {
	mtx := Matrix{Shape{height, width}, nil, nil}
	mtx.SeedData = build2DArray(mtx.Shape)
	mtx.Build()
	return mtx
}

func (m *Matrix) Build() {
	m.SeedData = build2DArray(m.Shape)
}

func (m *Matrix) Smooth(iterations int) {
	m.Smoothed = smooth2DArray(m.SeedData, iterations)
}

func build2DArray(Shape Shape) [][]int {
	var result [][]int
	for i := 0; i < Shape.height; i++ {
		var temp []int
		for j := 0; j < Shape.width; j++ {
			temp = append(temp, rand.Intn(20))
		}
		result = append(result, temp)
	}
	return result
}

func smooth2DArray(data [][]int, iterations int) [][]int {
	var result [][]int
	for i := 0; i < len(data); i++ {
		var temp []int
		for j := 0; j < len(data[i]); j++ {
			neighbors := []int{
				// account for modulo
				data[(i-1+len(data))%len(data)][(j-1+len(data[i]))%len(data[i])],
				data[(i-1+len(data))%len(data)][j],
				data[(i-1+len(data))%len(data)][(j+1)%len(data[i])],
				data[i][(j-1+len(data[i]))%len(data[i])],
				data[i][(j+1)%len(data[i])],
				data[(i+1)%len(data)][(j-1+len(data[i]))%len(data[i])],
				data[(i+1)%len(data)][j],
				data[(i+1)%len(data)][(j+1)%len(data[i])],
				data[i][j],
			}
			var sum int
			for _, v := range neighbors {
				sum += v
			}
			temp = append(temp, sum/len(neighbors))
		}
		result = append(result, temp)
	}
	if iterations > 1 {
		return smooth2DArray(result, iterations-1)
	}

	return result
}

func Display2DArray(data [][]int) {
	colorMap := buildColorMap(9, 15)
	for k, v := range colorMap {
		fmt.Printf("%d: %s\n", k, fmt.Sprintf("%v", v))
	}
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data[i]); j++ {
			var padded string
			if data[i][j] < 10 {
				padded = fmt.Sprintf("0%d", data[i][j])
			} else {
				padded = fmt.Sprintf("%d", data[i][j])
			}
			color.BlueString("%s", padded)
		}
		fmt.Println()
	}
}

type colorFunc func(string)

func buildColorMap(min int, max int) map[int]colorFunc {
	bluePercent := 0.5
	redPercent := 0.5
	blueRange := int(float64(max-min) * bluePercent)
	redRange := int(float64(max-min) * redPercent)
	blueStart := min
	blueEnd := blueStart + blueRange
	redStart := max - redRange
	redEnd := max

	colorMap := make(map[int]colorFunc)
	for i := blueStart; i < blueEnd; i++ {
		colorMap[i] = func(s string) { color.BlueString("%d", s) }
	}
	for i := blueEnd; i < redStart; i++ {
		colorMap[i] = func(s string) { color.BlueString("%d", s) }
	}
	for i := redStart; i < redEnd; i++ {
		colorMap[i] = func(s string) { color.BlueString("%d", s) }
	}
	return colorMap
}
