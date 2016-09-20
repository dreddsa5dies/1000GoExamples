// Дано целое число N (> 0). Используя один цикл, найти сумму
// 1 + 1/(1!) + 1/(2!) + 1/(3!) + ... + 1/( N !)
// (выражение N ! — N–факториал — обозначает произведение всех целых
// чисел от 1 до N : N ! = 1·2·...· N ). Полученное число является приближен-
// ным значением константы e = exp(1)

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	var (
		n int
	)

	n = -1
	for n < 0 {
		n = ioutil.Integer("N")
	}
	fmt.Println("1 + 1/(1!) + 1/(2!) + 1/(3!) + ... + 1/( N !)")
	fmt.Printf("%v\n", foo1(n))
}

func foo1(n int) float64 {
	k := 1.0
	for i := 1; i <= n; i++ {
		k = k + 1/foo2(i)
	}
	return k
}

func foo2(n int) float64 {
	k := 1.0
	for i := 1; i <= n; i++ {
		k = k * float64(i)
	}
	return k
}
