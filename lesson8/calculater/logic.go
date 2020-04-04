package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func process() error {
	for {
		fmt.Println()
		fmt.Println("please input expression:")
		input, _ := getInput()
		if input == "xxx\r\n" {
			fmt.Println("quit program")
			break
		}
		exSlice, express, err := preprocess(input)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("expression is:", express)
		calNodes, sufExpress, err := transToSuffixExpress(exSlice)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("transform to suffix-expression:", sufExpress)
		result, err := calculate(calNodes)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println("result of calculation:", result)
	}

	return nil
}

func getInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	return reader.ReadString('\n')
}

func preprocess(str string) (ret []string, express string, err error) {
	/*去除空格：检查有无错误字符*/
	s := make([]byte, 0, len(str)+1)
	for _, v := range str {
		if v == ' ' || v == '\r' || v == '\n' {
			continue
		}
		if (v >= '0' && v <= '9') || v == '.' || v == '(' || v == ')' {
			s = append(s, byte(v))
		} else if v == '+' || v == '-' || v == '*' || v == '/' || v == '^' {
			s = append(s, byte(v))
		} else {
			err = fmt.Errorf("input error")
			return nil, "", err
		}
	}
	express = string(s)
	s = append(s, '#') //添加结束符,便于后续处理结尾字符
	str = string(s)

	ret = make([]string, 0, len(s)-1)

	/*把表达式切分到字符串切片*/
	temp := str[:] //映射表达式
	j := 0         //记录数的首下标
	for i, v := range str {
		/*由于数的字符是连续的，如果当前不是数的字符，
		则表示这个数的字符已经遍历完，写入这个数到容器*/
		if !((v >= '0' && v <= '9') || v == '.') {
			temp = str[j:i]    //映射数的所有字符
			if len(temp) > 0 { //若得到有效的数
				ret = append(ret, temp) //写入容器
			}
			j = i + 1 //更新下个数的首下标
			if v != '#' {
				ret = append(ret, string(v)) //写入每个操作符,除了结束符
			}
		}
	}

	return
}

func transToSuffixExpress(s []string) (ret []Node, sufEx string, err error) {
	/*分配内存*/
	ret = make([]Node, 0, len(s))
	stack := new(StackStr)
	temp := new(Node)

	/*具体转换步骤*/
	for _, v := range s {
		if v[0] >= '0' && v[0] <= '9' {
			temp.num, _ = strconv.ParseFloat(v, 64)
			temp.isNum = true
			ret = append(ret, *temp)
		} else if v[0] == '+' || v[0] == '-' || v[0] == '*' || v[0] == '/' || v[0] == '^' {
			if stack.top == 0 {
				stack.push(v)
				continue
			}
			for t, _ := stack.getTop(); opPriority[v[0]] <= opPriority[t[0]]; t, _ = stack.getTop() {
				a, _ := stack.pop()
				temp.op = a[0]
				temp.isNum = false
				ret = append(ret, *temp)

				if stack.top == 0 {
					break
				}
			}
			stack.push(v)
		} else if v[0] == ')' {
			if stack.top == 0 {
				err = fmt.Errorf("express error")
				return
			}
			for t, _ := stack.getTop(); t[0] != '('; t, _ = stack.getTop() {
				a, _ := stack.pop()
				temp.op = a[0]
				temp.isNum = false
				ret = append(ret, *temp)

				if stack.top == 0 {
					err = fmt.Errorf("express error")
					return
				}
			}
			stack.pop()
		} else if v[0] == '(' {
			stack.push(v)
		}
	}

	/*按序输出栈内剩余操作符*/
	for stack.top != 0 {
		a, _ := stack.pop()
		if a[0] == '(' || a[0] == ')' {
			err = fmt.Errorf("express error")
			return
		}
		temp.op = a[0]
		temp.isNum = false
		ret = append(ret, *temp)
	}

	/*转化成字符串*/
	for i, v := range ret {
		if v.isNum == true {
			sufEx += fmt.Sprintf("%g", v.num)
		} else {
			sufEx += string(v.op)
		}
		if i < len(ret)-1 {
			sufEx += " "
		}
	}

	return
}

func calculate(nodes []Node) (result float64, err error) {
	stack := new(StackF64)
	/*遍历表达式*/
	for _, v := range nodes {
		if v.isNum == true { //当前节点是数则入栈
			stack.push(v.num)
		} else { //当前节点是操作符则进行计算

			/*如果栈内元素少于2个就计算则说明表输入的表达式有错误*/
			if stack.top < 2 {
				err = fmt.Errorf("express is incorrect")
				return
			}

			/*弹出栈顶两个元素进行计算*/
			a2, _ := stack.pop()
			a1, _ := stack.pop()
			switch v.op {
			case '+':
				a1 += a2
			case '-':
				a1 -= a2
			case '*':
				a1 *= a2
			case '/':
				a1 /= a2
			case '^':
				a1 = math.Pow(a1, a2)
			default: //无法识别的操作符
				err = fmt.Errorf("operator is unrecognized")
				return
			}
			stack.push(a1) //计算结果入栈
		}
	}

	/*遍历完表达式，如果栈内元素多于1个则说明输入的表达式不完整*/
	if stack.top > 1 {
		err = fmt.Errorf("express is incomplete")
		return
	}

	/*无错误，输出结果*/
	result, _ = stack.pop()
	return
}
