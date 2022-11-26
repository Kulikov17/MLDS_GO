package main

import (
	"fmt"
	"math"
)

func minDivisor(num int) int {
	var sqrtN = int(math.Sqrt(float64(num)))

	for i := 2; i <= sqrtN; i++ {
		if num%i == 0 {
			return i
		}
	}

	return num
}

func main() {

	var num int

	fmt.Scanf("%d", &num)

	fmt.Printf("%d", minDivisor(num))

}
