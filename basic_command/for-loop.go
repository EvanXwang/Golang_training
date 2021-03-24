package main

import "fmt"

func main() {
	/*
	var x int = 0
	for x < 3 {
		fmt.Println(x)
		x++
	}
	 */
 /*
	var x int
	for x=0;x<=10;x++{
		fmt.Println(x)
	}

  */
	var sum int =0
	var x int = 1
	for x <= 50{ //測試 1+2+3+50 的結果
		sum += x
		x++
	}
	fmt.Println(sum)
}