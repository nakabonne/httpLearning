package main

import "fmt"

func main() {
	method := standardIn()
	switch method {
	case 1:
		query()
	case 2:
		cookie()
	case 3:
		proxy()
	default:
		fmt.Println("数字が違います")
		main()
	}
}

func standardIn() int {
	fmt.Println("メソッドを選択して下さい")
	fmt.Println("1 単純なGetメソッド\n2 クッキー送信\n3 プロキシ送信")
	var i int
	fmt.Scan(&i)
	return i
}
