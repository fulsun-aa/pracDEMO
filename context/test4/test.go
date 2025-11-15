package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	deadline := time.Now().Add(5 * time.Second)
	ctx, _ := context.WithDeadline(context.Background(), deadline)
	go Speak(ctx)
	time.Sleep(time.Second * 10)

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
