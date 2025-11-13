package main

import (
	"fmt"
	"runtime"
	_ "time"
)

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	ch := make(chan int, 2)
	//我们还可以将cpu的核心数改为1，这样就只有一个goroutine在运行
	runtime.GOMAXPROCS(1)
	go func() {
		sum := 0
		for i := 0; i < len(slice)/2; i++ {
			sum += slice[i]
		}
		ch <- sum
	}()
	go func() {
		sum := 0
		for i := len(slice) / 2; i < len(slice); i++ {
			sum += slice[i]
		}
		ch <- sum
	}()
	//不能使用len(ch)，因为len(ch)一开始为空，放入一个就为1，取出一个就为0，随时变化，所以达不到遍历2次的效果
	for i := 0; i < 2; i++ {
		sum := 0
		for j := 0; j < 2; j++ {
			sum += <-ch
		}
	}
	fmt.Println(<-ch)

}
