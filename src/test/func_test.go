package test

// 测试需要满足条件
// 1. 文件以 _test.go结尾
// 2. 方法以Test开头
import (
	"fmt"
	"log"
	"runtime"
	"testing"
)

type bingOp func(a, b int) int

func TestFuncT(t *testing.T) {
	var res = func(a, b int) int {
		return a + b
	}(2, 8)
	fmt.Println(res)
}

func TestCaller(t *testing.T) {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d", file, line)
	}
	where()
}

func TestFibonacci(t *testing.T) {
	result := 0
	for i := 0; i < 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) = %d\n", i, result)
	}
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}
