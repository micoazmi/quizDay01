package main

import "fmt"

func main() {
	// findDivisor(6)
	findDivisor(24)
	// findDivisor(7)
}

func findDivisor(n int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			fmt.Printf("%d ", i)
		}

	}
}
