package main

import (
	"fmt"
	"sync"
	"time"
)

type obj struct {
}

var once sync.Once
var ob *obj

func getObj() *obj {
	once.Do(func() {
		//从打印结果可以看出来，尽管调用了一百个协程去执行getObj，但是只有一个协程会执行到create obj
		fmt.Println("create obj")
		ob = &obj{}
	})
	return ob
}

func main() {
	for i := 0; i < 100; i++ {
		go func() {

			_ = getObj()
		}()
	}
	time.Sleep(time.Second)
}
