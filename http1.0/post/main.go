package main

import "fmt"

func main() {
	method := standardIn()
	switch method {
	case 1:
		multipart_form_data()
	default:
		fmt.Println("数字が違います")
		main()
	}
}

func standardIn() int {
	fmt.Println("メソッドを選択して下さい")
	fmt.Println("1 multipart_form_data")
	var i int
	fmt.Scan(&i)
	return i
}
