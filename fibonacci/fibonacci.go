package main

import "fmt"

func fibDP(n int) (nthFib int) {
	var intArr []int
	intArr = append(intArr, 0)
	intArr = append(intArr, 1)
	if n > 1 {
		for i := 2; i <= n; i++ {
			intArr = append(intArr, (intArr[i-1] + intArr[i-2]))
		}

		nthFib = intArr[n]
	} else {
		nthFib = 0
	}

	return
}

func main() {
	fmt.Println(fibDP(1))
	fmt.Println(fibDP(5))
}
