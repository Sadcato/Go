package main

import (
	"Oct_30/test"
)

func main() {
	u := &test.User{
		Name: "张三",
		Age:  18,
	}
	s := &test.Student{
		Name: "李四",
		Age:  20,
	}
	s.Run()
	test.PersonCase(s)
	test.PersonCase(u)
	test.PersonCase2(u)
}
