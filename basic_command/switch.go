package main

import (
	"fmt"
	"time"
)

func main() {
	answer := 1
	switch answer {
	case 1:
		fmt.Println("答案為A")
	case 2:
		fmt.Println("答案為B")
	case 3:
		fmt.Println("答案為C")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("今天是假日，可以去看電影囉")
	default:
		fmt.Println("今天不是假日，繼續學程式")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12 :
		fmt.Println("目前為上午時段：健身")
	default:
		fmt.Println("目前為下午時段：跑步")
	}
}
