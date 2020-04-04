package main

import "fmt"

func strReverseV1(str string) {
	var strSlice []byte
	strSlice = []byte(str)

	for i := 0; i < len(strSlice)/2; i++ {
		temp := strSlice[i]
		strSlice[i] = strSlice[len(strSlice)-i-1]
		strSlice[len(strSlice)-i-1] = temp
	}

	str = string(strSlice)
	fmt.Println(str)
}

func strReverseV2(str string) {
	var strSlice []rune = []rune(str)

	for i := 0; i < len(strSlice)/2; i++ {
		temp := strSlice[i]
		strSlice[i] = strSlice[len(strSlice)-i-1]
		strSlice[len(strSlice)-i-1] = temp
	}

	str = string(strSlice)
	fmt.Println(str)
}

func strIsHuiWen(str string) {
	var strSlice []rune = []rune(str)
	var flag bool
	for i := 0; i < len(strSlice)/2; i++ {
		if strSlice[i] != strSlice[len(strSlice)-i-1] {
			flag = true
		}
	}
	if flag == false {
		fmt.Println(str, "是回文")
	} else {
		fmt.Println(str, "不是回文")
	}
}

func main() {
	strReverseV1("hello world")
	strReverseV1("hello 中国")
	strReverseV2("hello world")
	strReverseV2("hello 中国")
	strIsHuiWen("上海自来水来自海上")
	strIsHuiWen("发射点发生98w9df")
}
