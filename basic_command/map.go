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
	fmt.Println(m1["年紀"])

	fmt.Println(m1["身高"]) // 如果不存在這個key ，拿到對應值類型的零值
	value , ok := m1 ["身高"] //查詢map是否有值
	if !ok {
		fmt.Println("查無此key")
	}else {
		fmt.Println(value)
	}


	for k, v := range m1 {
		fmt.Println(k,v)
	}
}