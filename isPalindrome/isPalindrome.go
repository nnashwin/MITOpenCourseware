package main

import "fmt"

func isPalindrome(str string) bool {
	mid := len(str) / 2
	last := len(str) - 1
	for i := 0; i < mid; i++ {
		if str[i] != str[last-i] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(isPalindrome("cookies"))
	fmt.Println(isPalindrome("racecar"))
	fmt.Println(isPalindrome("madam"))
}
