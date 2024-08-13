package test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestFmt(t *testing.T) {
	fmt.Printf("%d + %d = %d\n", 1, 2, 3)
	var str = fmt.Sprintf("hello world! %s, time: %d, %d", "Tory", time.Now().UnixMilli(), time.Now().Unix())
	fmt.Println(str)

	fmt.Printf("name: %s, age: %d, height: %.2f\n", "john", 23, 89.995)
}

func TestNum(t *testing.T) {
	var a = 1.62
	var b = int(a)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(math.Ceil(a))
	fmt.Println(math.Floor(a))
	fmt.Println(math.Round(a))
	fmt.Println(rand.Int31n(20))
}
