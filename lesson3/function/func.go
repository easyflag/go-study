package main

import "fmt"

func isPrimeNumber(num int) bool {
	if num < 2 {
		return false
	}

	result := true
	for i := 2; i <= num/2; i++ {
		if num%i == 0 {
			result = false
			break
		}
	}
	return result
}

func printPrimeNumber(num int) {
	fmt.Printf("%d以内的质数有：", num)
	for i := 0; i <= num; i++ {
		if isPrimeNumber(i) == true {
			fmt.Printf("%d、", i)
		}
	}
	fmt.Printf("\n")
}

func printShuiXianHuaNumber() {
	for i := 101; i < 1000; i++ {
		sigalBit := i % 100 % 10
		tenBit := i % 100 / 10
		hundredBit := i / 100
		if i == hundredBit*hundredBit*hundredBit+tenBit*tenBit*tenBit+
			sigalBit*sigalBit*sigalBit {
			fmt.Printf("%d是水仙花数\n", i)
		}
	}
}

func main() {
	printPrimeNumber(1001)
	printShuiXianHuaNumber()
}
