package main

//Implement the dining philosopher’s problem with the following constraints/modifications.

//There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.

//Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture)

//The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).

//In order to eat, a philosopher must get permission from a host which executes in its own goroutine.

//The host allows no more than 2 philosophers to eat concurrently.

//Each philosopher is numbered, 1 through 5.

//When a philosopher starts eating (after it has obtained necessary locks) it prints “starting to eat <number>” on a line by itself, where <number> is the number of the philosopher.

//When a philosopher finishes eating (before it has released its locks) it prints “finishing eating <number>” on a line by itself, where <number> is the number of the philosopher.

import (
	"log"
	"sync"
	"time"
)

const MaxPhilos = 5
const MaxMeals = 3
const MaxEaters = 2

type chopS struct {
	mut sync.Mutex
}

type host struct {
	activeEaters []*philosopher
}

func (h *host) isPermitted(p *philosopher) bool {
	// TODO:
	// Add eater if there is enough space and is the eater is not currently eating
	return false
}

type philosopher struct {
	id              int
	mealNum         int
	leftCs, rightCs *chopS
}

func (p *philosopher) eat(h *host, wg *sync.WaitGroup) {
	defer wg.Done()
	if h.isPermitted(p) {
		for p.mealNum < MaxMeals {
			p.leftCs.mut.Lock()
			p.rightCs.mut.Lock()

			log.Printf("id: %d starting to eat", p.id)
			time.Sleep(time.Second)
			log.Printf("id: %d eating", p.id)
			time.Sleep(time.Second)
			log.Printf("id: %d finishing to eat", p.id)

			p.mealNum += 1
			p.rightCs.mut.Unlock()
			p.leftCs.mut.Unlock()
		}
	}
}

func main() {
	cSticks := make([]*chopS, MaxPhilos)
	philos := make([]*philosopher, MaxPhilos)
	var wg sync.WaitGroup

	for i := 0; i < MaxPhilos; i++ {
		cSticks[i] = new(chopS)
	}

	for i := 0; i < MaxPhilos-1; i++ {
		philos[i] = &philosopher{
			id:      i + 1,
			mealNum: 0,
			leftCs:  cSticks[i],
			rightCs: cSticks[i+1],
		}
	}

	philos[len(philos)-1] = &philosopher{
		id:      MaxPhilos,
		mealNum: 0,
		leftCs:  cSticks[0],
		rightCs: cSticks[MaxPhilos-1],
	}

	h := host{activeEaters: make([]*philosopher, MaxEaters)}

	wg.Add(MaxPhilos)
	for i := 0; i < MaxPhilos; i++ {
		go philos[i].eat(&h, &wg)
	}
	wg.Wait()
}
