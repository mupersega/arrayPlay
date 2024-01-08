package main

import (
	"fmt"
	"math/rand"

	"github.com/fatih/color"
)

type shape struct {
	height int
	width  int
}

type matrix struct {
	shape    shape
	seedData [][]int
	smoothed [][]int
}

func main() {
	m := matrix{shape{10, 10}, nil, nil}
	m.build()
	m.smooth()

	display2DArray(m.seedData)
	fmt.Println()
	display2DArray(m.smoothed)
	aMap := buildColorMap(0, 20)
	for k, v := range aMap {
		fmt.Printf("%d: %s\n", k, v)
	}
}

func (m *matrix) build() {
	m.seedData = build2DArray(m.shape)
}

func (m *matrix) smooth() {
	m.smoothed = smooth2DArray(m.seedData, 1)
}

func build2DArray(shape shape) [][]int {
	var result [][]int
	for i := 0; i < shape.height; i++ {
		var temp []int
		for j := 0; j < shape.width; j++ {
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

func display2DArray(data [][]int) {
	colorMap := buildColorMap(9, 15)
	for k, v := range colorMap {
		fmt.Printf("%d: %s\n", k, v)
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
	// redPercent := 0.5
	blueRange := int(float64(max-min) * bluePercent)
	// redRange := int(float64(max-min) * redPercent)
	blueStart := min
	blueEnd := blueStart + blueRange
	// redStart := max - redRange
	// redEnd := max

	colorMap := make(map[int]colorFunc)
	for i := blueStart; i < blueEnd; i++ {
		colorMap[i] = func(s string) { color.BlueString("%d", s) }
	}
	// for i := blueEnd; i < redStart; i++ {
	// 	colorMap[i] = "white"
	// }
	// for i := redStart; i < redEnd; i++ {
	// 	colorMap[i] = "red"
	// }
	return colorMap
}
