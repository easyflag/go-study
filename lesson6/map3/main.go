package main

import (
	"fmt"
)

func main() {
	mapSlice := make([]map[string]int, 2, 4)
	mapSlice[0] = make(map[string]int)
	fmt.Printf("len:%d  cap:%d", len(mapSlice), cap(mapSlice))
	mapSlice[0]["s1"] = 99
	mapSlice[0]["s2"] = 79

	sliceMap := make(map[string][]int)
	sliceMap["s1"] = make([]int, 2, 4)
	sliceMap["s1"][0] = 99

}
