package main

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

func main() {
	fmt.Printf("プロセスID: %d\n", os.Getpid())
	fmt.Printf("親プロセスID: %d\n", os.Getppid())
	sid, _ := syscall.Getsid(os.Getpid())
	fmt.Fprintf(os.Stderr, "グループID: %d セッションID: %d\n",
		syscall.Getpgrp(), sid)

	// psコマンドでプロセス確認用
	time.Sleep(10 * time.Second)
	fmt.Printf("sleep end")
}
