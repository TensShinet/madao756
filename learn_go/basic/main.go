package main

// 使用 goroutine
func goroutineMain() {
	go Say("Hi")
	Say("ts")
}

func main() {
	goroutineMain()
}
