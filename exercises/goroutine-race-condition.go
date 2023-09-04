package exercices

import (
	"fmt"
	"sync"
)

func ExecuteGoRoutineRaceCondition() {
	var wg sync.WaitGroup
	numberRoutines := 100
	counter := 0

	wg.Add(numberRoutines)
	for i := 0; i < numberRoutines; i++ {
		go func () {
			defer wg.Done()
			for i := 0; i < 100; i++ {
				counter++
			}
		}()
	}
	wg.Wait()

	fmt.Println("Final Counter:", counter)
}

/*
A race condition occurs in a concurrent or multi-threaded program when
two or more threads or processes access a shared resource simultaneously,
and at least one of them attempts to modify that resource. The program's
behavior becomes dependent on the exact timing of these concurrent accesses,
potentially leading to unexpected and unintended outcomes. In the example code above,
even using WaitGroup to wait for all goroutines to complete, the output leads to 
unpredictable and incorrect results. This race condition occurs because there
is no synchronization mechanism (such as mutexes) in place to ensure exclusive
access to the counter variable.
*/
