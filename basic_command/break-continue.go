package main

import "fmt"

func main() {
	/*
	var x int = 0
	for x < 10{
		if x==5{
			break
		}
		fmt.Println(x)
		x++
	}
	 */
	//實際範例，持續讓使用者輸入數字，計算總和，直到使用者輸入0為止
	var sum int = 0
	for true { //無窮迴圈 最好搭配break
		var n int
		fmt.Scan(&n)
		sum +=n

		if n ==0{
			break
		}

	}
	fmt.Println("目前總和為：",sum)
}