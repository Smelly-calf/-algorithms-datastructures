package main

import (
	"fmt"
	"strconv"
)

// 整数的四则运算
// 第一步：中缀表达式变后缀表达式: exp-> postexp，9,3,1,-,3,*,10,2,/,+,+
// 第二步：后缀表达式利用一个数字栈计算结果

func main() {
	ret, _ := calc([]string{"9", "3", "1", "-", "3", "*", "11", "2", "/", "+", "+"})
	fmt.Println(ret)
}

func calc(postFixExpress []string) (result int, err error) {
	var (
		num1 string
		num2 string
		s    Stack
	)

	for _, post := range postFixExpress {
		//fmt.Printf("post:%s\n", post)
		if post == "+" || post == "-" || post == "*" || post == "/" {
			num1, err = s.Pop()
			if err != nil {
				return
			}
			num2, err = s.Pop()
			if err != nil {
				return
			}
			//fmt.Printf("num1:%s, num2:%s\n", num1, num2)

			B, _ := strconv.Atoi(num1)
			A, _ := strconv.Atoi(num2)
			var ret int
			switch post {
			case "+":
				ret = A + B
			case "-":
				ret = A - B
			case "*":
				ret = A * B
			case "/":
				ret = A / B
			default:
				panic("非法字符")
			}
			s.Push(strconv.Itoa(ret))
		} else {
			s.Push(post)
		}
		//fmt.Printf("%+v\n", s.data[:s.top+1])
	}
	result2, err := s.Top()
	if err != nil {
		return
	}
	result, err = strconv.Atoi(result2)
	return
}

type Stack struct {
	data [1024]string
	top  int
}

func (s *Stack) Push(st string) {
	s.data[s.top] = st
	s.top++
}

func (s *Stack) Top() (ret string, err error) {
	if s.top == 0 {
		err = fmt.Errorf("栈空")
		return
	}
	return s.data[s.top-1], nil
}

func (s *Stack) Pop() (ret string, err error) {
	if s.top == 0 {
		err = fmt.Errorf("stack is empty")
		return
	}
	s.top--
	ret = s.data[s.top]
	return
}
