package main

import "sync"

var mutex sync.Mutex

func buyTicket(wg *sync.WaitGroup, userId int, tickets *int) {
	mutex.Lock()
	defer wg.Done()
	if *tickets > 0 {
		*tickets--
		println("User", userId, "bought a ticket")
	} else {
		println("Tickets are sold out")
	}
	mutex.Unlock()
}

func main() {
	var tickets int = 500

	var wg = sync.WaitGroup{}

	for userId := 0; userId <= 2000; userId++ {
		wg.Add(1)
		go buyTicket(&wg, userId, &tickets)
	}
	wg.Wait()
}
