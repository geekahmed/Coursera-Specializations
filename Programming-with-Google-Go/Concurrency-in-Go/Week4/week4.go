package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopSticks struct {
	sync.Mutex
}
type Philosopher struct {
	lChop, rChop *ChopSticks
	id int
}
var wg sync.WaitGroup

func (P Philosopher) eat(){

	for i := 0; i < 3 ; i++{
		P.lChop.Lock()
		P.rChop.Lock()

		PrintPhilosoper("starting to eat", P.id)
		time.Sleep(2 * time.Second)
		PrintPhilosoper("finishing eating", P.id)
		P.lChop.Unlock()
		P.rChop.Unlock()
		time.Sleep(2 * time.Second)
	}

	wg.Done()

}

func main() {
	chops := make([]*ChopSticks, 5)
	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		chops[i] = new(ChopSticks)
	}
	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{id: i, lChop: chops[i], rChop: chops[(i+1) % 5]}
		wg.Add(1)
		go philosophers[i].eat()
	}

	wg.Wait()

}

func PrintPhilosoper(name string, id int){
	fmt.Printf("#%d is %s\n", id+1, name)
}
