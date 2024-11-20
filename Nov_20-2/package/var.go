package _case

import "fmt"

func VarDeclareCase() {
	var i, z, c int = 1, 2, 3
	var j = 100
	var f float32 = 100.12
	B := true
	var arr = [4]int{1, 2, 3, 4}
	arr1 := [...]int{2, 3, 4, 5}
	var arr2 [3]int
	arr2[2] = 4
	arr2[1] = 4
	fmt.Print(i, z, c, j, f, B, arr1, arr2, arr)
}
