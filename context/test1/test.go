package main

import (
	"context"
	"fmt"
	"time"
)

var traceId = "id"
var value = "value"

func main() {
	ctx := context.WithValue(context.Background(), traceId, value)
	userctx := context.WithValue(ctx, "user", "user123")
	go getTraceId(userctx)
	time.Sleep(time.Second)
}
func getTraceId(ctx context.Context) {
	fmt.Println("key:", traceId, "value:", ctx.Value(traceId))
}
