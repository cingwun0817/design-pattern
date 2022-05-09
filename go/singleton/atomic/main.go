package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type single struct {
}

var singleInstance *single

var flag uint32

var lock = &sync.Mutex{}

/**
 * use `Mutex Lock` and `Atomic`
 */
func getInstance() *single {
	if atomic.LoadUint32(&flag) == 1 {
		fmt.Println("Single instance already created.")
		return singleInstance
	}

	lock.Lock()
	defer lock.Unlock()
	if flag == 0 {
		fmt.Println("Creating single instance now.")
		singleInstance = &single{}
		atomic.StoreUint32(&flag, 1)
	} else {
		fmt.Println("Single instance already created.")
	}

	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
