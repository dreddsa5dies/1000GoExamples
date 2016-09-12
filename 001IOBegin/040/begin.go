// Найти решение системы линейных уравнений вида
// A 1 · x + B 1 · y = C 1 ,
// A 2 · x + B 2 · y = C 2 ,
// заданной своими коэффициентами A 1 , B 1 , C 1 , A 2 , B 2 , C 2 , если известно, что
// данная система имеет единственное решение. Воспользоваться формулами
// x = ( C 1 · B 2 – C 2 · B 1 )/ D ,
// y = ( A 1 · C 2 – A 2 · C 1 )/ D ,
// где D = A 1 · B 2 – A 2 · B 1
package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var a1, a2, b1, b2, c1, c2, d, x, y float64
	fmt.Println("\tA1 · x + B1 · y = C1\n\tA2 · x + B2 · y = C2")
	a1 = ioutil.Number("Введите A1")
	b1 = ioutil.Number("Введите B1")
	c1 = ioutil.Number("Введите C1")
	a2 = ioutil.Number("Введите A2")
	b2 = ioutil.Number("Введите B2")
	c2 = ioutil.Number("Введите C2")
	d = a1*b2 - a2*b1
	x = (c1*b2 - c2*b1) / d
	y = (a1*c2 - a2*c1) / d
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}
