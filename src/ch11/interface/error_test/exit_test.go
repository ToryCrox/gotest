package error_test

import (
	"errors"
	"fmt"
	"testing"
)


func TestExit(t *testing.T){
	defer func(){
		if err := recover(); err != nil {
			fmt.Println("recoverd from", err)
		}
	}()
	fmt.Println("Start")
	panic(errors.New("Something wrong!"))
}