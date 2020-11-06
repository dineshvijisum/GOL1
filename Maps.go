package main

import "fmt"

func main() {
	celebs := map[string]int{
		"Nicolas Cage":       50,
		"Selena Gomez":       21,
		"Jude Law":           41,
		"Scarlett Johansson": 29,
		"Dinesh": 5,
	}
	d := map[float64]int{
		2.22:2,
		2333:2,
	}

	fmt.Printf("%#v", celebs)
	fmt.Printf("%#v",d)
}