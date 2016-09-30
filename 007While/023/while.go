// Даны целые положительные числа A и B . Найти их наибольший об-
// щий делитель (НОД), используя алгоритм Евклида :
// НОД( A , B ) = НОД( B , A mod B ), если B ≠ 0;
// НОД( A , 0) = A ,
// где «mod» обозначает операцию взятия остатка от деления

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		a, b int
	)

	for a <= 0 {
		a = ioutil.Integer("A")
	}
	for b <= 0 {
		b = ioutil.Integer("B")
	}
	x := countLoop(a, b)
	fmt.Printf("НОД A и B = ")
	fmt.Printf("%v\n", x)
}

func countLoop(a, b int) int {
	for a != 0 && b != 0 {
		if a > b {
			a = a % b
		} else {
			b = b % a
		}
	}
	return a + b
}
