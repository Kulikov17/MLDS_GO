package main

import (
	"fmt"
	"math"
)

func main() {

	var num1, num2, num3, num4, res float64

	fmt.Scanf("%f\n", &num1)
	fmt.Scanf("%f\n", &num2)
	fmt.Scanf("%f\n", &num3)
	fmt.Scanf("%f\n", &num4)

	res = math.Min(num1, math.Min(num2, math.Min(num3, num4)))

	fmt.Printf("%d", int(res))

}
