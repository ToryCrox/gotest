package share_mem_test

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	var mut sync.Mutex
	var wg sync.WaitGroup
	counter := 0
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock()
				wg.Done()
			}()
			mut.Lock()
			counter++

		}()
	}
	wg.Wait()
	t.Log(counter)
}

func service() string {
	time.Sleep(time.Millisecond * 1050)
	return "service Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 1500)
	fmt.Println("Task is done")
}

//channel
func AsyncService() chan string {
	//第二个参数表示channel的buffer
	retCh := make(chan string, 1)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret //放到channel中，如果有buffer，则继续执行，否则等待接收
		fmt.Println("service exited.")
	}()
	return retCh
}

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Millisecond * 20)
}

func TestSelect(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("time Error")
	}
}

func dataProducer(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataReceiver(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for {
			if i, ok := <-ch; ok {
				fmt.Println(i)
			} else {
				break
			}
		}
		wg.Done()
	}()
}

func TestPR(t *testing.T) {
	var wg sync.WaitGroup
	ch := make(chan int, 5)
	wg.Add(1)
	dataProducer(ch, &wg)
	wg.Add(1)
	dataReceiver(ch, &wg)
	wg.Wait()

}

func isCancelled(ctx context.Context) bool{
	select{
	case <- ctx.Done():
		return true
	default:
		return false
	}
}

func TestCancel(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context){
			for {
				if isCancelled(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 50)
			}
			fmt.Println(i, "Cancelled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
