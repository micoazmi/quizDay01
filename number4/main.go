package main

import "fmt"

func main() {
	piramid()

}

func piramid() {
	var input int
	fmt.Print("input Jumlah baris piramid: ")
	fmt.Scan(&input)
	for row := 0; row < input; row++ {
		for col := input; col > row; col-- {
			fmt.Printf("%d ", col)
		}
		fmt.Println(" ")
	}
}
