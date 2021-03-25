package main

import "fmt"

// n個台階，一次可以走一步，也可以走二步，有多少種走法
// 遞歸實現如下：

func step (n uint64)uint64 {
	if n == 1 {
		//如果只有一個台階就只有一種走法
		return 1
	}
	if n == 2 {
		return 2
	}
	return step(n-1) + step(n-2)
}

func main() {
	ret := step(10)
	fmt.Println(ret)

}