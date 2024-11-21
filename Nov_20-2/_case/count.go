package _case

import "fmt"

func ConstAnd() {
	const (
		a = 10
		b = 20
	)
	fmt.Println(a, b)
}
