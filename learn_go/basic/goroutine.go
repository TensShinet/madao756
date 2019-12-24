package main

import (
	"fmt"
	"runtime"
)

func Say(s string) {
	for i := 0; i < 5; i++ {
		// 时间片让给别人，也就是自己不执行了，给另一个 goroutine
		runtime.Gosched()
		fmt.Println(s)
	}
}

// 使用 goroutine
func goroutineMain() {
	go Say("Hi")
	Say("ts")
}
