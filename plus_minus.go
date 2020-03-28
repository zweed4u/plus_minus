package main

import "fmt"

func plusMinus(ar []int) []float64 {
	positiveNums := 0
	negativeNums := 0
	zeroNums := 0
	denom := len(ar)
	for _, number := range ar {
		if number == 0 {
			zeroNums += 1
		} else if number < 0 {
			negativeNums += 1
		} else {
			positiveNums += 1
		}
	}

	positiveFrac := float64(positiveNums) / float64(denom)
	fmt.Println(positiveFrac)

	negativeFrac := float64(negativeNums) / float64(denom)
	fmt.Println(negativeFrac)

	zeroFrac := float64(zeroNums) / float64(denom)
	fmt.Println(zeroFrac)
	return []float64{positiveFrac, negativeFrac, zeroFrac}
}

func main() {
	arr := []int{1, 2, 3, 0, 0, -1}
	plusMinus(arr)
}
