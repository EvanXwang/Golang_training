package main

import "fmt"

func main() {

	/*
	var x int
	fmt.Print("請輸入你的成績：")
	fmt.Scanln(&x)
	if x >= 60 && x <= 100 {
		fmt.Println("做得好")
	} else if  x <= 60 && x>= 50 {
		fmt.Println("加把勁")
	} else {
		fmt.Println("你是在哈囉")
	}
	*/

	var money int
	fmt.Println("請問想領多少錢??")
	fmt.Scanln(&money)
	if money < 50 {
		fmt.Println("太少了")
	}else if money < 100000 {
		fmt.Println("OK")
	}else {
		fmt.Println("too much")
	}
	fmt.Println("執行完畢")
}