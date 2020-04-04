package main

import (
	"fmt"
	"time"
)

func insertSort(a [10]int) [10]int {
	len := len(a)
	for i := 1; i < len; i++ {
		for j := i; j > 0; j-- {
			if a[j] < a[j-1] {
				a[j], a[j-1] = a[j-1], a[j]
			} else {
				break
			}
		}
	}
	return a
}

func selectSort(a [10]int) [10]int {
	len := len(a)
	for i := 0; i < len-1; i++ {
		for j := i; j < len-1; j++ {
			if a[j+1] < a[i] {
				a[j+1], a[i] = a[i], a[j+1]
			}
		}

	}
	return a
}

func floatSort(a [10]int) [10]int {
	len := len(a)
	for i := 0; i < len-1; i++ {
		for j := 0; j < len-1-i; j++ {
			if a[j+1] < a[j] {
				a[j+1], a[j] = a[j], a[j+1]
			}
		}
	}
	return a
}

func mySortV1(a [10]int) [10]int {
	len := len(a)
	for i := 0; i < len; i++ {
		index_tp := 0
		for j := 1; j < len-i; j++ {
			if a[j] > a[index_tp] {
				index_tp = j
			}
		}
		temp := a[index_tp]
		for j := index_tp; j < len-1; j++ {
			a[j] = a[j+1]
		}
		a[len-1] = temp
	}
	return a
}

func mySortV2(a [10]int) [10]int {
	len := len(a)
	for i := 0; i < len-1; i++ {
		index_tp := 0
		for j := 1; j < len-i; j++ {
			if a[j] > a[index_tp] {
				index_tp = j
			}
		}
		a[index_tp], a[len-1-i] = a[len-1-i], a[index_tp]
	}
	return a
}

func timer1(a [10]int) {
	start := time.Now().Second()
	for i := 0; i < 100000000; i++ {
		_ = insertSort(a)
	}
	end := time.Now().Second()
	timeConsume := (end - start)
	fmt.Printf("insertSort耗时：%ds\n", timeConsume)
}

func timer2(a [10]int) {
	start := time.Now().Second()
	for i := 0; i < 100000000; i++ {
		_ = selectSort(a)
	}
	end := time.Now().Second()
	timeConsume := (end - start)
	fmt.Printf("selectSort耗时：%ds\n", timeConsume)
}

func timer3(a [10]int) {
	start := time.Now().Second()
	for i := 0; i < 100000000; i++ {
		_ = floatSort(a)
	}
	end := time.Now().Second()
	timeConsume := (end - start)
	fmt.Printf("floatSort耗时：%ds\n", timeConsume)
}

func timer4(a [10]int) {
	start := time.Now().Second()
	for i := 0; i < 100000000; i++ {
		_ = mySortV1(a)
	}
	end := time.Now().Second()
	timeConsume := (end - start)
	fmt.Printf("mySortV1耗时：%ds\n", timeConsume)
}

func timer5(a [10]int) {
	start := time.Now().Second()
	for i := 0; i < 100000000; i++ {
		_ = mySortV2(a)
	}
	end := time.Now().Second()
	timeConsume := (end - start)
	fmt.Printf("mySortV2耗时：%ds\n", timeConsume)
}

func main() {
	var a [10]int = [10]int{55, 12, 31, 56, 18, 34, 67, 99, 77, 66}
	//var b [10]int
	//for i := 0; i < 10; i++ {
	//	fmt.Scanf("%d\n", &a[i])
	//}
	fmt.Println(a)
	//b := insertSort(a)
	//b := selectSort(a)
	//b := floatSort(a)
	//b := mySortV1(a)
	//b := mySortV2(a)

	go timer1(a)
	go timer2(a)
	go timer3(a)
	go timer4(a)
	go timer5(a)

	time.Sleep(time.Second * 15)
	fmt.Println("over")
}
