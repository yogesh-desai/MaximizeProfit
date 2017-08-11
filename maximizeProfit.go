package main

import (

	"fmt"
)

func main(){

	var input = []int{6, 5, 4, 7, 8, 5, 9, 7, 6}
	min, max, diff := getMaxDiff(input)
	fmt.Println("Index of smallest number: ", min)
	fmt.Println("Index of largest number: ", max)
	fmt.Println("Maximum Difference: ", diff)
}

func getMaxDiff(input []int) (int, int, int) {

	min := input[0] //Assume first integer as smallest
	diff := input[1] - input[0] // assume first two intergers have maximum difference
	var m, l int //To store indexes of the smallest and largest numbers.

	for i := 0; i < len(input); i++ {

		if (input[i] - min) > diff { diff = input[i] - min ; l = i }

		if input[i] < min { min = input[i]; m = i  }
		
	}
	return m,l,diff
}

/* OUTPUT

Index of smallest number:  2
Index of largest number:  6
Maximum Difference:  5

*/
