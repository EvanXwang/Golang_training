package main

import "fmt"

func main() {
	/*
	show ("hello")
	show("你好")
	 */

	//呼叫函式，取得回傳值
	/*
	var x int  = multiply(3,3)
	fmt.Println(x)
	 */

	//
	/*
	var d int = add(3,4)
	fmt.Println(d)
	 */

	//多個回傳值列印
	var a int
	var b string

	a,b = test()
	fmt.Println(a,b)

}


//執行return , 結束函式
func show(msg string){
	if msg =="hello"{
		return
	}
	fmt.Println(msg)
}

//也可以return 一個數值，但需定義型態
func add(n1 int, n2 int) int{
	var result int = n1 + n2
	fmt.Println(result)
	return 10
}


//可以做整數乘法的函式，大刮號的int 為 return的參數值型態
func multiply (n1 int, n2 int)int{
	var result int = n1 * n2
	return result
}

//可傳 多個回傳值
func test ()(int, string){
	return 1, "hello"

}

