package main

import (
	"fmt"
)

func replaceAtIndex(str string, replacement rune, index int) string {
	return str[:index] + string(replacement) + str[index+1:]
}

func genSkittles(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += "I"
	}
	return s
}

func main() {

	var n, k, l, r int
	var res string

	fmt.Scanf("%d %d", &n, &k)

	res = genSkittles(n)

	for i := 0; i < k; i++ {
		fmt.Scanf("%d %d", &l, &r)

		for j := l; j <= r; j++ {
			res = replaceAtIndex(res, '.', j-1)
		}
	}

	fmt.Printf("%s", res)

}
