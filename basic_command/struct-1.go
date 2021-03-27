package main

import "fmt"

//結構體模擬實現其他語言中的"繼承"
type animal struct {
	name string
}

//給animal實現一個移動的方法
func (a animal) move (){
	fmt.Printf("%s會動",a.name)
}


//狗
type dog struct {
	feet uint8
	animal   //animal 擁有的方法，dog 此時也有了繼承
}

//給狗實現一個汪的方法
func (d dog) wang (){
	fmt.Printf("%s在叫：汪汪汪", d.name)
}

func main(){
	d1 := dog {
		animal:animal{name:"jojo"},
		feet:  4,
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()


}