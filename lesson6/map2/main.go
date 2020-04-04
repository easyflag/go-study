package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	a := make(map[string]int)
	for i := 0; i < 100; i++ {
		key := fmt.Sprintf("key%d", i)
		a[key] = rand.Intn(1000)
	}

	for key, val := range a {
		fmt.Printf("a[%s] = %d\n", key, val)
	}

	keySlice := make([]string, 0, len(a))
	for i := range a {
		keySlice = append(keySlice, i)
	}
	for i, v := range keySlice {
		fmt.Printf("keySlice[%d] = %s\n", i, v)
	}
	sort.Strings(keySlice)
	for i, v := range keySlice {
		fmt.Printf("keySlice[%d] = %s\n", i, v)
	}

	for _, v := range keySlice {
		fmt.Printf("a[%s] = %d\n", v, a[v])
	}
}
