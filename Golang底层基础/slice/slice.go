package main

import "fmt"

func main()  {
	slice := make([]int,2)
	_ = append(slice,1,2,3)
	fmt.Println(slice)
}
