package main

//测试一下gosched
import (
	"fmt"
	"runtime"
	_ "time"
)

func main() {
	//我们还可以将cpu的核心数改为1，这样就只有一个goroutine在运行
	runtime.GOMAXPROCS(1)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("goroutine", i)
		}
	}()
	//让出cpu时间片，让其他线程执行，但也有可能一个时间片不够其他线程执行完毕
	//执行结果为：有时为goroutine0，有时没有打印，说明狗语言会调用多个核心的cpu来执行。
	//但是我们将cpu的核心数改为1，打印结果为
	// goroutine 0
	// goroutine 1
	// goroutine 2
	// goroutine 3
	// goroutine 4
	// goroutine 5
	// goroutine 6
	// goroutine 7
	// goroutine 8
	// goroutine 9
	runtime.Gosched()
}
