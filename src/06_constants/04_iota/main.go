package main

import "fmt"

const (
	_ = iota //0
	B = iota*10 //10
	A = iota*10 // 20
)

func main()  {
	fmt.Println(B)
	fmt.Println(A)
}
