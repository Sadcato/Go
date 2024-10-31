package test

import "fmt"

type User struct {
	Name string
	Age  int
}

func (*User) Run() {
	fmt.Println("user run")

}
func (*User) Sleep() {
	fmt.Println("user run")

}

type Person interface {
	Run()
	Sleep()
}
type Student struct {
	Name string
	Age  int
}

func (*Student) Run() {
	fmt.Println("student run")
}
func (*Student) Sleep() {
	fmt.Println("student sleep")
}
func PersonCase(person Person) {
	person.Run()
	person.Sleep()

}
func PersonCase2(person interface{}) {
	p1, ok := person.(Person)
	if ok {
		p1.Run()
	} else {
		fmt.Println("not a person")
	}
}
