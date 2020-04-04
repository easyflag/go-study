package main

import (
	"fmt"
)

func main() {
	a1 := make(map[string]int, 2)
	fmt.Printf("len:%d\n", len(a1))
	a1["stu1"] = 100
	a1["stu3"] = 67
	a1["stu2"] = 77
	delete(a1, "stu1")
	for key, val := range a1 {
		fmt.Printf("a1[%s] = %d\n", key, val)
	}
	fmt.Printf("len:%d\n", len(a1))
	fmt.Printf("%v\n", a1)

	a2 := make([]int, 2, 2)
	fmt.Printf("type of a1 is %T\n", a1)
	fmt.Printf("len:%d  cap:%d\n", len(a2), cap(a2))
	a2[0] = 75
	a2[1] = 76
	a2 = append(a2, 45)
	fmt.Printf("len:%d  cap:%d\n", len(a2), cap(a2))
	fmt.Printf("%v\n", a2)
}
