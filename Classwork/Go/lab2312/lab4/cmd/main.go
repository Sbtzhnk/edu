package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func f(v *uint32, wg *sync.WaitGroup) {
	for i := 0; i < 3000; i++ {
		atomic.AddUint32(v, 1)
	}
	wg.Done()
}

func main() {
	var v uint32 = 44
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)
	go f(&v, &waitGroup)
	go f(&v, &waitGroup)
	waitGroup.Wait()

	fmt.Println(v)
}
