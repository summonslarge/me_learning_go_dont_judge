package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	nums := []int{5, 6, 7, 8, 9}

	for _, val := range nums {
		fmt.Println(val)
	}

}
