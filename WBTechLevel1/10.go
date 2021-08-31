package main

import (
	"fmt"
	"math"
)

// Дана последовательность температурных колебаний (-25.4, -27.0 13.0, 19.0, 15.5, 24.5, -21.0, 32.5). Объединить данный значения в группы с шагов в 10 градусов.
// Последовательность в подмножностве не важна.
// Пример: (-20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc)

func main() {
	var arr = []float64{-29.9, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, 9, 10, 1000, 1001, -5, -9, -11, 9999, 9989}
	result := make(map[float64][]float64)
	for idx := range arr {
		val := math.Ceil(arr[idx]) // округление вниз
		valRem := int(val) % 10
		result[val-float64(valRem)] = append(result[val-float64(valRem)], arr[idx])
	}

	fmt.Println(result)
}
