// By definition, a race condition is a condition of a program where its behavior depends on relative timing
// or interleaving of multiple threads or processes.
// One or more possible outcomes may be undesirable, resulting in a bug.
// We refer to this kind of behavior as nondeterministic.

package main

import (
	"fmt"
	"sync"
)

func inc(p *int, wg *sync.WaitGroup) {
	defer wg.Done()
	*p++
}

func printVal(p *int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100000; i++ {
	}
	fmt.Printf("value = %d\n", *p)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	p := 0

	go inc(&p, &wg)
	go printVal(&p, &wg)

	wg.Wait()
}
