package exercices

import (
	"fmt"
	"sync"
)

func ExecuteGoRoutineRaceCondition() {
	var wg sync.WaitGroup
	// var mx sync.Mutex // solution
	numberRoutines := 100
	counter := 0

	wg.Add(numberRoutines)
	for i := 0; i < numberRoutines; i++ {
		go func () {
			defer wg.Done()
			// mx.Lock()
			for i := 0; i < 100; i++ {
				counter++
			}
			// mx.Unlock()
		}()
	}
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

