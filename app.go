package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ch1 := make(chan int, 95)
	wg.Add(1)
	go func() {
		for i := range ch1 {
			fmt.Println(i)
			time.Sleep(time.Second * 1)
		}
		wg.Done()
	}()
	for i := range 100 {
		ch1 <- i
		fmt.Println("Length: ", len(ch1))
	}
	fmt.Println("All values sent.")
	wg.Wait()
}
