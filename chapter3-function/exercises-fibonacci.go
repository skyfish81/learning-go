package main

import "fmt"

func main() {
	// fmt.Println(fibonacci(1))
	for _, term := range fibonacci(10) {
		fmt.Println(term)
	}
}

func fibonacci(value int) []int {
	x := make([]int, value)
	x[0], x[1] = 1, 1
	for n := 2; n < value; n++ {
		x[n] = x[n-1] + x[n-2]
	}
	return x
	// if n == 1 {
	// 	return 1
	// } else {
	// 	return n + fibonacci(n-1)
	// }
}
