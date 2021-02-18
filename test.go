package main

import "fmt"

func magic(i int) (ret int := 2) {
	if i > 12 {
		ret = i
	}
	return
}

func main() {
	fmt.Printf("should be 2: %d\n", magic(0))
	fmt.Printf("should be 27: %d\n", magic(27))
}
