package main

import "fmt"

func zero(x int)  {
	fmt.Printf("%p\n",&x) // address in func zero
	fmt.Println(&x) // address in func zero
	x = 0
}


func zeroref(x *int)  {
	fmt.Printf("%p ref\n",&x) // address in func zero
	fmt.Println(&x) // address in func zero
	*x = 0
}

func main()  {
	x:= 5
	fmt.Printf("%p \n", &x)  // address in main
	fmt.Println(&x) // address in mian
	zero(x)
	fmt.Println(x) // x still 5

	zeroref(&x)
	fmt.Println(x) // x still 5
}


