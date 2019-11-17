package main

import (
	"io"
	"os"
)

func main() {
	file, err := os.Open("file.go")
	if err != nil {
		panic(err)
	}
	// deferを利用するとメソッド内の全ての処理が完了後にその処理を行う。
	// 今回の場合はファイルを閉じる処理がio.Copyの後に行われる
	defer file.Close()
	io.Copy(os.Stdout, file)
}
