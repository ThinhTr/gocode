package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main()  {
	res, _ := http.Get("http://www.mcleods.com/")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)


}
