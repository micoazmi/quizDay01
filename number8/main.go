package main

import (
	"fmt"
	"strings"
)

func main() {
	testCase1 := isPalindrome("Kasur ini rusak")
	testCase2 := isPalindrome("tamaT")
	testCase3 := isPalindrome("Aku Usa")

	fmt.Println(testCase1)
	fmt.Println(testCase2)
	fmt.Println(testCase3)

}

func isPalindrome(word string) bool {
	reverseWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		reverseWord += string(word[i])
	}

	if strings.ToLower(reverseWord) == strings.ToLower(word) {
		return true
	} else {
		return false
	}

}
