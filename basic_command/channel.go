package main

import "fmt"

func sum ( s [] int , c chan int ) {
	sum := 0
	for _ , v := range s {
		sum += v
	}
	c < - sum //把sum發送到通道c
}

func main () {
	s := [] int { 7 , 2 , 8 , - 9 , 4 , 0 }

	c := make ( chan int )
	go sum ( s [: len ( s ) / 2 ], c )
	go sum ( s [ len ( s ) / 2 :], c )
	x , y := < - c , < - c //從通道c中接收

	fmt.Println ( x , y , x + y )
}
