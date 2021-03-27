package main

import (
	"fmt"
	"os"
)

/*
函數版學生管理系統
寫一個系統能夠查看\新增\刪除 學生
 */

var (
	allStudent map[int64]*student  //變量聲明
)

type student struct {
	id   int64
	name string
}

//newStudent 是student 類型的構造函數
func newStudent(id int64,name string) *student{
	return &student{
		id: id,
		name : name,
	}
}


func show(){

	//把所有的學生都列印出來
	for k, v:= range allStudent {
		fmt.Printf("學號：%d 姓名：%s\n",k,v.name)
	}
}

func add(){
	fmt.Println("新增學生")
	//向allStudent中新增一個新的學生
	//1.新增一個學生
	//1.1 獲取用戶輸入
	var (
		id int64
		name string
	)
	fmt.Println("請輸入學生學號：")
	fmt.Scanln(&id)
	fmt.Println("請輸入學生姓名：")
	fmt.Scanln(&name)
	//1.2 新增學生 (調用student 的構造函數)
	newStu := newStudent(id , name)
	//2 追加到allStudent 這個map中
	allStudent[id] = newStu
}

func del(){
	// 1.請用戶輸入要刪除的學生的學號
	var (
		deleteId int64
	)
	fmt.Print("請輸入學生學號：")
	fmt.Scanln(&deleteId)
	// 2.去allStudent 這個map中根據學號刪除對應的鍵值
	delete(allStudent,deleteId)
}




func main() {
	allStudent = make(map[int64]*student,48)//初始化 (開啟內存空間)
	// 1 列印選單
	for {

		fmt.Println("Welcome to Student System!!")
		fmt.Println(
			"\n1.查看所有學生\n2.新增學生\n3.刪除學生\n4.退出")

		fmt.Println("請輸入執行編號：")
		// 2 等待用戶選擇要做什麼
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你選擇了%d這個選項!\n", choice)
		// 3 執行對應的函數
		switch choice {
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
}

