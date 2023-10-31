package main

import (
	"fmt"
	"sync"
)

type single struct {
}

var (
	once           sync.Once
	singleInstance *single
)

func getInstance() *single {
	once.Do(func() {
		fmt.Println("Creating single instance now.")
		singleInstance = &single{}
	})
	fmt.Println("Returning single instance.")
	return singleInstance
}

func main() {
	for i := 0; i < 30; i++ {
		go getInstance()
	}

	_, _ = fmt.Scanln()
}
