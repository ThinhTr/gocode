package main

import "fmt"

const(
	_ =iota
	KB = 1<<(iota*10)
	MB= 1<<(iota*10) // 2*10
	GB= 1<<(iota*10)
	TB= 1<<(iota*10)
)

func main()  {
	fmt.Println("binary\t\tdecimal")
	fmt.Printf("%b\t", KB)
	fmt.Printf("%d\t",KB)
	fmt.Printf("%b\t",MB)
	fmt.Printf("%d\t",MB)
	fmt.Printf("%b\t", GB)
	fmt.Printf("%d\t",GB)
	fmt.Printf("%b\t",TB)
	fmt.Printf("b\t",TB)

}