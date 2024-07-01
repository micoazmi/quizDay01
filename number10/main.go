package main

import "fmt"

func main() {
	testCase1 := checkBraces("(())")
	testCase2 := checkBraces("()()")
	testCase3 := checkBraces("((()")
	testCase4 := checkBraces("(()))((())")
	fmt.Println(testCase1)
	fmt.Println(testCase2)
	fmt.Println(testCase3)
	fmt.Println(testCase4)

}

func checkBraces(brace string) bool {
	counterLeft := 0
	counterRight := 0
	for i := 0; i < len(brace); i++ {
		if brace[i] == '(' {
			counterLeft++
		} else {
			counterRight++
		}
	}
	if counterLeft == counterRight {
		return true
	} else {
		return false
	}
}
