package main

import (
	"fmt"
	"math"
)

func main() {

	var n float64

	fmt.Scanf("%f\n", &n)

	sqrt := int(math.Sqrt(n))

	for i := 0; i <= sqrt; i++ {
		for j := i; j <= sqrt; j++ {
			for k := j; k <= sqrt; k++ {
				for l := k; l <= sqrt; l++ {
					if (i*i + j*j + k*k + l*l) == int(n) {

						if (i != 0) {
							fmt.Printf("%d %d %d %d", i, j, k, l)
							return
						}

						if (j != 0) {
							fmt.Printf("%d %d %d", j, k, l)
							return
						}

						if (k != 0) {
							fmt.Printf("%d %d", k, l)
							return
						}
						
						fmt.Printf("%d", l)
						return
					}
				}
			}
		}
	}
}
