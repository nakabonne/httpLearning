package main

import "fmt"

func main() {
	method := standardIn()
	switch method {
	case 1:
		simpleGet()
	case 2:
		cookie()
	default:
		fmt.Println("数字が違います")
		main()
	}
}

func standardIn() int {
	fmt.Println("メソッドを選択して下さい")
	fmt.Println("1 単純なGetメソッド\n2 クッキー送信")
	var i int
	fmt.Scan(&i)
	return i
}
