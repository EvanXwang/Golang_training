package main

import "fmt"

//map 是一種無序的基於 key - value 的數據結構
//定義語法  map [keyType] ValueType


func main() {
	var m1 map [string] int
	// 需初始化 (沒有在內存中新增空間)
	m1 = make (map [string ]int , 10) // 要先計算好該map容量，避免在程式在運行期間，再動態擴容
	m1["年紀"] = 18
	m1["體重"] = 60
	fmt.Println(m1)
}