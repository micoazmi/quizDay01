package main

import "fmt"

func main() {
	angkaDeret(9)
	angkaDeret(5)

}

func angkaDeret(n int) {
	number := n
	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			if col%2 == 0 {
				fmt.Printf("%d ", number)
			} else {
				fmt.Printf("%d ", row)
			}
		}
		number--
		fmt.Println(" ")
	}
}
