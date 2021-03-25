package main

import "fmt"
//3! = 3*2*1     = 3*2!
//4! = 4*3*2*1   = 4*3!
//5! = 5*4*3*2*1 = 5*4!

//遞歸：函數自己調用自己
//適合處理 問題相同，問題規模越來越小的場景
//遞歸一定要有一個明確的退出條件
//計算n的階乘
func f(n uint64) uint64{
	if n <=1 {
		return 1
	}
	return n * f(n-1)
}

func main() {

	ret:= f(30)
	fmt.Println(ret)
}
