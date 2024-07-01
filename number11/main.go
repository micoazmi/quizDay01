package main

import "fmt"

func main() {
	testCase1 := isNumberPalindrome(121)
	fmt.Println(testCase1)

	testCase2 := isNumberPalindrome((2147447412))
	fmt.Println(testCase2)

	testCase3 := isNumberPalindrome((333))
	fmt.Println(testCase3)

	testCase4 := isNumberPalindrome((110))
	fmt.Println(testCase4)

	testCase5 := isNumberPalindrome((11))
	fmt.Println(testCase5)
}

func isNumberPalindrome(n int) bool {
	if n < 0 {
		return false
	}
	original, reversed := n, 0
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10

	}
	return original == reversed

}
