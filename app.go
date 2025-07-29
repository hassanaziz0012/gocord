package main

import (
	"gocord/web"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		web.StartHttp()
		wg.Done()
	}()

	wg.Wait()
}
