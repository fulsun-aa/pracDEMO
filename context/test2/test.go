package main

import (
	"context"
	"fmt"
	"time"
)

/*
select的 “立即判断” 还是 “等待”，取决于是否有default分支。
具体来说：
select的执行逻辑是先一次性检查所有case中的通道操作是否 “就绪”（即：是否可以立即执行，不阻塞）：
如果有至少一个case就绪（比如通道可接收 / 发送数据，或已关闭），select会随机选一个就绪的case执行，整个过程是 “立即” 的（不等待）。
如果所有case都未就绪（都会阻塞）：
若有default分支，select会立即执行default（不等待任何case）；
若没有default分支，select会阻塞当前 goroutine，直到某个case变为就绪（此时再执行该case）。
*/
func main() {
	//返回一个cancelCtx上下文和一个取消函数
	ctx, cancel := context.WithCancel(context.Background())
	go Speak(ctx)
	//等待3秒
	time.Sleep(time.Second * 3)
	//取消上下文，会向ctx.Done()的通道发送一个值
	cancel()
	//等待1秒，确保 Speak 函数中的 ctx.Done() 被触发
	time.Sleep(time.Second)

}
func Speak(ctx context.Context) {
	//time.Tick的功能是按照每个参数时间发送一条数据，它是一个通道，返回值为chan time.Time
	for range time.Tick(time.Second) {
		select {
		case <-ctx.Done():
			fmt.Println("stop speaking")
			return
		default:
			fmt.Println("keep speaking")
		}
	}
}
