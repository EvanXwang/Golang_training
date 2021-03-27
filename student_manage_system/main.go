package main

import (
	"fmt"
	"os"
)

/*
函數版學生管理系統
寫一個系統能夠查看\新增\刪除 學生
 */
func show(){
	fmt.Println("查看所有學生")
}

func add(){
	fmt.Println("新增學生")
}

func del(){
	fmt.Println("刪除學生")
}




func main() {
	// 1 列印選單
	fmt.Println("Welcome to Student System!!")
	fmt.Println(
		"\n1.查看所有學生\n2.新增學生\n3.刪除學生\n4.退出")

	fmt.Println("請輸入執行編號：")
	// 2 等待用戶選擇要做什麼
	var choice int
	fmt.Scanln(&choice)
	fmt.Printf("你選擇了%d這個選項!\n",choice)
	// 3 執行對應的函數
	switch choice{
	case 1:
		show()
	case 2:
		add()
	case 3:
		del()
	case 4:
		os.Exit(1)
	default:
		fmt.Println("重新輸入")
	}
}

