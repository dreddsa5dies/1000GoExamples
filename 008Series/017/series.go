// Дано вещественное число B , целое число N и набор из N веществен-
// ных чисел, упорядоченных по возрастанию. Вывести элементы набора
// вместе с числом B , сохраняя упорядоченность выводимых чисел

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	b := ioutil.Number("число B")
	n := ioutil.Integer("число N")

	count := series(n)
	fmt.Printf("несортированныей набор:\t%v\n", count)
	sort(count, 0, len(count)-1)
	fmt.Printf("сортированныей набор:\t%v\n", count)
	insert(b, count)
}

func series(n int) []int {
	var (
		r   int
		num []int
	)
	for i := 0; i < n; i++ {
		r = ioutil.Integer("число")
		num = append(num, r)
	}
	return num
}

func sort(num []int, start, end int) {
	if start >= end {
		return
	}

	pivot := num[start]
	i := start + 1

	for j := start; j <= end; j++ {
		if pivot > num[j] {
			num[i], num[j] = num[j], num[i]
			i++
		}
	}

	num[start], num[i-1] = num[i-1], num[start]

	sort(num, start, i-2)
	sort(num, i, end)
}

func insert(b float64, count []int) {
	fmt.Printf("Вставка: ")
	for i := 0; i < len(count); i++ {
		switch {
		case count[i] == int(b):
			fmt.Printf("| %v ", b)
		case float64(count[i]) < b:
			fmt.Printf("| %v ", count[i])
		case float64(count[i]) > b:
			fmt.Printf("| %v ", count[i])
		}
	}
	fmt.Println()
}
