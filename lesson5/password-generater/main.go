package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

var (
	length int
	form   string
)

const (
	numStr  string = "0123456789"
	charStr string = "AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	specStr string = "`~!@#$%^&*()-_=+\\|[]{};:'\",.<>/?"
)

func parseArgs() {
	flag.IntVar(&length, "l", 64, "-l:指定生成的密码长度")
	flag.StringVar(&form, "t", "advance",
		`-t:指定生成的密码组合形式
	 	 num：只使用数字
		 mix：使用数字和字母
		 advance：使用数字、字母、特殊字符`)
	flag.Parse()
}

func generatePassword() string {
	password := make([]byte, length, length)
	var souceStr string

	if form == "num" {
		souceStr = fmt.Sprintf("%s", numStr)
	} else if form == "mix" {
		souceStr = fmt.Sprintf("%s%s", numStr, charStr)
	} else if form == "advance" {
		souceStr = fmt.Sprintf("%s%s%s", numStr, charStr, specStr)
	} else {
		souceStr = fmt.Sprintf("%s", numStr)
	}

	tempStr := []byte(souceStr)
	for i := 0; i < length*length*length; i++ {
		x := rand.Intn(length)
		y := rand.Intn(length)
		tempStr[x], tempStr[y] = tempStr[y], tempStr[x]
	}

	for i := 0; i < length; i++ {
		password[i] = tempStr[rand.Intn(length)]
	}

	return string(password)

}

func main() {
	rand.Seed(time.Now().UnixNano())
	parseArgs()
	fmt.Printf("length:%d  form:%s\n", length, form)
	password := generatePassword()
	fmt.Println(password)
}
