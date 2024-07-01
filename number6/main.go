package main

import "fmt"

func main() {
	angkaDeret(9)
	fmt.Println(" ")
	angkaDeret(5)

}

func angkaDeret(n int) {
	for row := 1; row <= n; row++ {
		for col := 1; col <= n; col++ {
			if col%2 == 0 && row%2 == 0 {
				fmt.Printf("%s ", "-")

			} else if col%2 == 0 || row%2 == 0 {

				fmt.Printf("%d ", col)
			} else {
				fmt.Printf("%s ", "-")

			}
		}
		fmt.Println(" ")
	}
}
