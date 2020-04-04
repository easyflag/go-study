package main

import (
	"fmt"
	"math/rand"
	"time"
)

func v1(array [100]int, sum int) {
	for i := 0; i < len(array)-1; i++ {
		for j := i + 1; j < len(array); j++ {
			if array[i]+array[j] == sum {
				fmt.Printf("%d(array[%d]) + %d(array[%d]) = %d\n", array[i], i, array[j], j, sum)
			}
		}
	}
}

func main() {
	var a [100]int

	rand.Seed(time.Now().Unix())
	for i := range a {
		a[i] = rand.Intn(100)
	}

	v1(a, 137)
}
