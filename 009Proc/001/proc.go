// Описать процедуру PowerA3( A , B ), вычисляющую третью степень числа
// A и возвращающую ее в переменной B ( A — входной, B — выходной пара-
// метр; оба параметра являются вещественными). С помощью этой процеду-
// ры найти третьи степени пяти данных чисел.

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var b float64
	n := ioutil.Integer("Количество чисел")
	arr := ioutil.RandomFlt(n)
	for i := 0; i < len(arr); i++ {
		b = powerA3(arr[i])
		fmt.Printf("третья степень числа %v = %v\n", arr[i], b)
	}
}

func powerA3(a float64) (b float64) {
	b = math.Pow(a, 3)
	return b
}
