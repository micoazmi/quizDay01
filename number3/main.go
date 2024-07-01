package main

import "fmt"

func main() {
	segitiga(5)
	segitiga2(5)
	// segitiga3(5)
}

func segitiga(n int) {

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if col < row {
				fmt.Printf("%s ", " ")

			} else {

				fmt.Printf("%s ", "*")
			}

		}
		fmt.Println(" ")

	}

}

func segitiga2(n int) {

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if col == 4 {

				fmt.Printf("%s ", "*")
			} else if col == 3 && row != 0 {
				fmt.Printf("%s ", "*")

			} else if col == 2 && row != 0 && row != 1 {
				fmt.Printf("%s ", "*")

			} else if col == 1 && row != 0 && row != 1 && row != 2 {
				fmt.Printf("%s ", "*")

			} else if col == 0 && row != 0 && row != 1 && row != 2 && row != 3 {
				fmt.Printf("%s ", "*")

			} else {
				fmt.Printf("%s ", " ")
			}

		}
		fmt.Println(" ")

	}

}

func segitiga3(n int) {

	for row := 0; row < n; row++ {
		for col := 0; col < n; col++ {
			if col == 4 {
				fmt.Printf("%s ", "*")
			} else if col == 3 && row != 0 {
				fmt.Printf("%s ", "*")
			} else {
				fmt.Printf("%s ", " ")

			}
		}
		fmt.Println(" ")

	}

}
