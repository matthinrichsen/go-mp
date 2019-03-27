package main

import (
	"log"
	"sync"
)

func main() {
	c := make(chan int32, 9)
	consumerThreads := &sync.WaitGroup{}
	consumerThreads.Add(9)

	lastRounds := []Number{}
	newOnes := make([][]Number, 9)

	record := int64(1)
	recordLock := &sync.Mutex{}

	for i := 0; i < 9; i++ {
		lastRounds = append(lastRounds, FromInt64(int64(i)))

		go func() {
			for {
				number := <-c

				for _, n := range lastRounds {
					n = n.Add(number + 1)

					mp := n.MP()

					recordLock.Lock()
					if mp > record {
						log.Printf("%d: %d", n.ToInt(), mp)
						record = mp
					}
					recordLock.Unlock()

					newOnes[int(number)] = append(newOnes[int(number)], n)
				}

				consumerThreads.Done()
			}
		}()
	}

	for iters := 1; ; iters++ {
		log.Println(iters)
		for i := int32(0); i < 9; i++ {
			c <- i
		}

		consumerThreads.Wait()

		lastRounds = make([]Number, 0, 9*len(lastRounds))
		for _, n := range newOnes {
			lastRounds = append(lastRounds, n...)
		}

		consumerThreads.Add(9)
	}
}
