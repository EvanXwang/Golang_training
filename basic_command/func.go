package main

import "fmt"

//可以印出任何訊息的funtion
func showHello(msg string){
	fmt.Println("hello",msg)
}


func print(msg string){
	fmt.Println(msg)
}

//可以做整數乘法的funtion
//兩個以上的參數用逗號隔開
func sum(n1 int, n2 int){
	var result = n1+n2
	fmt.Println(result)
}


//計算1+2+3+...+10 的函式包裝
func sum_number(a int){
	var result int =0
	var n int
	for n=1;n<=a;n++{
		result +=n

	}
	fmt.Println(result)
}

//main函式 就是go程式的進入點
func main() {
	fmt.Println("請輸入一個數值：")
	var d int
	fmt.Scanln(&d)
	sum_number(d)

}
