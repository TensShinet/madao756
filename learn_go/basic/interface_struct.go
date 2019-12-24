package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

func (h *Human) SayHi(name string) {
	fmt.Printf("Hi %s, I am %s you can call me on %s\n", name, h.name, h.phone)
}

func (s *Student) SayHi(name string) {
	fmt.Printf("Hi %s, I am a student.My school is %s", name, s.school)
}

// 定义一组含有 SayHi 的接口
type HumanInterface interface {
	SayHi(name string)
}

func interfaceBasicMain() {
	h1 := Human{
		"tenshine",
		20,
		"888",
	}

	s1 := Student{
		Human{
			"zjt",
			20,
			"888",
		},
		"sxu",
	}

	h1.SayHi("gallon")
	s1.SayHi("ts")
}
