package exercices

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var host = make(chan bool, 2)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	id int
	leftCS, rightCS *ChopStick
}

func (p *Philosopher) eat() {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		host <- true
		
		p.leftCS.Lock()
		p.rightCS.Lock()
		fmt.Printf("[%d] starting to eat\n", p.id)
		time.Sleep(1 * time.Second)
		fmt.Printf("[%d] finished eating\n", p.id)
		p.leftCS.Unlock()
		p.rightCS.Unlock()

		<-host
	}

}

func ExecuteGoRoutineDiningPhiloshopers() {
	count := 5
	chopSticks := make([]*ChopStick, count)

	for i := range chopSticks {
		chopSticks[i] = new(ChopStick)
	}

	philosophers := make([]*Philosopher, count)

	for i := range philosophers {
		philosophers[i] = &Philosopher{
			id: i+1,
			rightCS: chopSticks[(i + 1) % count],
			leftCS: chopSticks[i],
		}
	}

	for _, philosopher := range philosophers {
		wg.Add(1)
		go philosopher.eat()
	}
	wg.Wait()
}
