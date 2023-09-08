package exercices

import (
	"fmt"
	"sync"
	"time"
)

func benchmark(f func()) {
	startTime := time.Now()

	f()

	endTime := time.Now()
	elapsedTime := endTime.Sub(startTime)

	fmt.Println(elapsedTime)
}

func exec() {
	var wg sync.WaitGroup
	var mx sync.Mutex
	var once sync.Once
	numberRoutines := 100
	counter := 0

	wg.Add(numberRoutines)
	for i := 0; i < numberRoutines; i++ {
		go func () {
			defer wg.Done()

			once.Do(func() {
				fmt.Println("oi")
			})

			mx.Lock()
			for i := 0; i < 100; i++ {
				counter++
			}
			mx.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

func ExecuteGoRoutinesBenchmark() {
	benchmark(exec)
}
