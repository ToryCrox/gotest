package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("hello, world\ntime is %v\n", time.Now().UnixMilli())
	fmt.Printf("hello, world\ntime is %v\n", time.Now().UnixNano())
	fmt.Printf("max: %v", max(100, 99))
	var str = fmt.Sprintf("hello, world\ntime is %v\n", time.Now().UnixMilli())
	fmt.Println(str)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("numbers: %v\n", numbers)

}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
