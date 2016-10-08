// Дано целое число N и набор из N целых чисел, упорядоченный по
// возрастанию. Данный набор может содержать одинаковые элементы. Вы-
// вести в том же порядке все различные элементы данного набора

package main

import (
	"fmt"

	"github.com/dreddsa5dies/1000GoExamples/ioutil"
)

func main() {
	n := ioutil.Integer("число N")

	count := series(n)
	fmt.Printf("%v\n", count)
	sort(count, 0, len(count)-1)
	fmt.Printf("%v\n", count)
	fmt.Printf("%v\n", delRepeat(count))
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

func delRepeat(count []int) []int {
	check := 0
	for i := 0; i < len(count)-1; i++ {
		if count[i] == count[i+1] {
			copy(count[i:], count[i+1:])
			count[len(count)-1] = 0
			check++
		}
	}
	return count[:(len(count) - check)]
}
