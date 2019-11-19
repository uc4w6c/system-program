package main

import (
	"context"
	"fmt"
)

func main() {
	fmt.Println("start sub()")
	// 終了を受け取るための終了関数つきコンテキスト
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		fmt.Println("sub() is finished")
		// 終了通知
		cancel()
	}()
	// 終了を待つ
	<-ctx.Done()
	fmt.Println("all tasks are finished")
}
