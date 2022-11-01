package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/panjf2000/ants/v2"
)

var sum int32

func sumAdd(i interface{}) {
	n := i.(int32)
	atomic.AddInt32(&sum, n)
	fmt.Printf("sumAdd, n = %+v, sum=%+v\n", n, sum)
}

func doSth() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("doSth,hello world")
}

func main() {
	defer ants.Release()

	runTimes := 1000

	var wg sync.WaitGroup
	work := func() {
		doSth()
		wg.Done()
	}
	for i := 0; i < runTimes; i++ {
		wg.Add(1)
		ants.Submit(work)
	}
	wg.Wait()

	select {}
}
