package main

import (
	"golang.org/x/exp/slices"
	"log"
	"sync"
)

const (
	PhilosNum = 5
	MealsNum  = 3
	MaxEaters = 2
)

type chopS struct {
	mut sync.Mutex
}

type host struct {
	activeEaters []*philosopher
}

func (h *host) host(askChan chan *philosopher, answChan chan *philosopher, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < PhilosNum*MealsNum; i++ {
		p := <-askChan
		if len(h.activeEaters) < MaxEaters {
			h.activeEaters = append(h.activeEaters, p)
			log.Printf("Allowed philosopher %d to eat\n", p.id)
			answChan <- p
			fullPhilo := <-answChan
			if slices.Contains(h.activeEaters, fullPhilo) {
				idx := slices.Index(h.activeEaters, fullPhilo)
				h.activeEaters = append(h.activeEaters[:idx], h.activeEaters[idx+1:]...)
			}
		}
	}
}

type philosopher struct {
	id              int
	mealNum         int
	leftCs, rightCs *chopS
}

func (p *philosopher) eat(askChan chan *philosopher, answChan chan *philosopher, wg *sync.WaitGroup) {
	defer wg.Done()
	for p.mealNum < MealsNum {
		askChan <- p
		if allowed := <-answChan; allowed != nil {
			p.leftCs.mut.Lock()
			p.rightCs.mut.Lock()
			log.Printf("Philosopher: %d starting to eat", p.id)
			log.Printf("Philosopher: %d eating", p.id)
			log.Printf("Philosopher: %d finishing to eat", p.id)
			p.mealNum += 1
			p.rightCs.mut.Unlock()
			p.leftCs.mut.Unlock()
			answChan <- p
		}
	}
}

func main() {
	cSticks := make([]*chopS, PhilosNum)
	philos := make([]*philosopher, PhilosNum)
	askChan := make(chan *philosopher)
	answChan := make(chan *philosopher)
	h := host{activeEaters: make([]*philosopher, 0)}
	var wg sync.WaitGroup

	for i := 0; i < PhilosNum; i++ {
		cSticks[i] = new(chopS)
	}
	for i := 0; i < PhilosNum; i++ {
		philos[i] = &philosopher{
			id:      i + 1,
			mealNum: 0,
			leftCs:  cSticks[i],
			rightCs: cSticks[(i+1)%5],
		}
	}

	wg.Add(PhilosNum + 1)
	go h.host(askChan, answChan, &wg)
	for i := 0; i < PhilosNum; i++ {
		go philos[i].eat(askChan, answChan, &wg)
	}
	wg.Wait()
}
