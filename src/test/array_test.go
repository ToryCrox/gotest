package test

import (
	"fmt"
	"testing"
)

func TestArray1(t *testing.T) {
	var arr1 [5]int
	fmt.Println("arr1:", arr1)
	fmt.Printf("len %d cap %d \n", len(arr1), cap(arr1))
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i + 1
	}
	fmt.Println("arr1:", arr1)
	for i, v := range arr1 {
		fmt.Printf("arr1[%d]=%d\n", i, v)
	}
	// 数组拷贝是值拷贝
	arr2 := arr1
	for i, v := range arr2 {
		arr2[i] = v * 2
	}
	fmt.Println("arr1:", arr1)
	fmt.Println("arr2:", arr2)
	arr3 := &arr1
	for i, v := range *arr3 {
		(*arr3)[i] = v * 3
	}
	fmt.Println("arr1:", arr1)
	fmt.Println("arr3:", *arr3)
}
