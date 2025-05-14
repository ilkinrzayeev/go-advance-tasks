package main

import (
	"fmt"
)

func countBits(num uint32) int32 {
	var count int32 = 0
	for num != 0 {
		count += int32(num & 1)
		num >>= 1
	}
	return count
}

func main() {
	var num uint32
	fmt.Scan(&num)
	result := countBits(num)
	fmt.Println(result)
}
