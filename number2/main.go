package main

import "fmt"

func main() {
	extractDigit(12234)
	// // extractDigit(5432)
	// extractDigit(1278)
}

func extractDigit(n int) {
	for n > 0 {
		extract := n % 10
		fmt.Printf("%d ", extract)
		n = n / 10
	}
}
