package main

import "fmt"

// 给函数实现方法，这个有一点抽象
// 函数的声明不能直接实现接口，需要将函数定义为类型后，使用类型实现结构体，当类型方法被调用时，还需要调用函数本体。

type SayFuck func(statement string)

func (f SayFuck) Say(statement string) {
	statement = "Fuck " + statement
	f(statement)
}

func Hi(statement string) {
	fmt.Println(statement)
}

func functionTypeMain() {
	statement := "ts"
	SayFuck(Hi).Say(statement)
}
