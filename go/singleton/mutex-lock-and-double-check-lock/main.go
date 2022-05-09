package main

import (
	"fmt"
	"sync"
)

type single struct {
}

var singleInstance *single

var lock = &sync.Mutex{}

/**
 * use `Mutex Lock` and `Double Check Lock`
 */
func getInstance() *single {
	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		if singleInstance == nil {
			fmt.Println("Creating single instance now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Single instance already created.")
		}
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
