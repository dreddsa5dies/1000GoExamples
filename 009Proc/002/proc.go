// Описать процедуру PowerA234( A , B , C , D ), вычисляющую вторую, тре-
// тью и четвертую степень числа A и возвращающую эти степени соответст-
// венно в переменных B , C и D ( A — входной, B , C , D — выходные парамет-
// ры; все параметры являются вещественными). С помощью этой процедуры
// найти вторую, третью и четвертую степень пяти данных чисел

package main

import (
	"fmt"
	"math"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("Количество чисел")
	arr := ioutil.RandomFlt(n)
	for i := 0; i < len(arr); i++ {
		b, c, d := powerA345(arr[i])
		fmt.Printf("%v^3 = %v\n%v^4 = %v\n%v^5 = %v\n", arr[i], b, arr[i], c, arr[i], d)
	}
}

func powerA345(a float64) (b, c, d float64) {
	b = math.Pow(a, 3)
	c = math.Pow(a, 4)
	d = math.Pow(a, 5)
	return b, c, d
}
