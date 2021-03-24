package main

import "fmt"

//常量
//定義了常量之後不能修改
//在程序運行期間不會改變的量
const pi = 3.1419
const e = 2.7182

//批量聲明常量
const (
	statusok = 200
	notfound = 400
)

//批量聲明常量，後面的值也會默認 上一行的參數
const (
	n1 = 100
	n2
	n3
)

//iota 是go語言的常量計數器
const (
	a1 = iota //0
	a2        //1
	a3        //2
)

//定義數量級
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {

	fmt.Println("a3:", KB)
}
