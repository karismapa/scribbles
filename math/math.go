package main

import (
	"fmt"
	"sync"
)

type cont struct {
	sync.RWMutex
	arr []int
}

func (c *cont) getArr() []int {
	c.RLock()
	defer c.RUnlock()

	return c.arr
}

func (c *cont) addAll() (sum int) {
	arr := c.getArr()

	if len(arr) > 0 {
		sum = arr[0]

		for _, num := range arr[1:] {
			sum += num
		}
	}

	return
}

func main() {
	// var addResult int
	// var subResult int
	// var mulResult int
	// var divResult int

	arr := []int{}
	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
	}

	cont1 := cont{arr: arr}

	sumResult := make(chan int)

	go func() {
		sumResult <- cont1.addAll()
	}()

	fmt.Println("Sum result:", <-sumResult)
}
