// Дан диаметр окружности d. Найти ее длину L = π·d. В качестве значения π использовать 3.14
package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/util"
)

func main() {
	var x float64
	fmt.Println("Необходимо найти длину окружности")
	x = util.Number("Введите значение диаметра окружности")
	fmt.Println("Длина равна:\t", x*math.Pi)
}
