package main

import "fmt"

func main()  {
	a:= 43
	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)    // This is call dereferencing   with the *


}
