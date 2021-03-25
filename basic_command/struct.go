package main

import "fmt"

//結構體

type person struct {
	 name 	string
	 age 	int
	 gender string
	 hobby 	[]string
}

func main() {
	// 聲明一個person類型的變量P
	var p person
	// 通過字段賦值
	p.name = "evan"
	p.age  = 37
	p.gender = "男"
	p.hobby =[]string{"電玩", "音樂", "電影"}
	fmt.Println(p)
	// 訪問變量p的字段
	fmt.Printf("%T\n",p)
	fmt.Println(p.name)
}