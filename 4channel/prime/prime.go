package main

import (
	"fmt"
	"math"
)

func primeNumber() chan int {
	result := make(chan int)
	go func() {
		result <- 2
		// // 値の返却タイミングを確かめる
		// for i := 3; i < 10; i += 2 {
		for i := 3; i < 100000; i += 2 {
			l := int(math.Sqrt(float64(i)))
			found := false
			for j := 3; j < l; j += 2 {
				if i%j == 0 {
					found = true
					break
				}
			}
			if !found {
				result <- i
			}
			// 値の返却タイミングを確かめる
			// time.Sleep(time.Second)
		}
		close(result)
	}()
	return result
}

func main() {
	pn := primeNumber()
	// ここがポイント
	// pnに全て値が入ってから処理するのではなく、処理ごとに以下に入る
	for n := range pn {
		fmt.Println(n)
	}
}
