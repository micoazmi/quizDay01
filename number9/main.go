package main

import (
	"fmt"
)

func main() {
	reverse("ABCD")
	reverse("tamaT")
	reverse("XYnb")
}

func reverse(word string) {

	for i := len(word) - 1; i >= 0; i-- {
		fmt.Printf("%v", string(word[i]))
	}

}
