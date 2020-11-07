package main

import "fmt"

func reverseBits(num uint32) uint32 {
	bit := make([]uint32, 32)
	for i := 0; i < 32; i++ {
		bit[i] = num & 1
		num = num >> 1
	}
	var res uint32
	for i := 0; i < 32; i++ {
		res = res*2 + bit[i]
	}
	return res
}

func main() {
	var num uint32 = 4294967293
	fmt.Println(reverseBits(num))
}
