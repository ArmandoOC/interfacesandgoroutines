package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func showGouroutine(id int, wg *sync.WaitGroup) {
	delay := rand.Intn(500)
	fmt.Printf("Goroutine %d with %d ms\n", id, delay)

	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done() //De esta manera la gouroutine avisa que ya terminó de ejecutarse.
}

func main() {
	start := time.Now()

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go showGouroutine(i, wg)
	}

	wg.Wait() //Hace que se espere hasta que todas las gouritines asociadas al waitgroup terminen de ejecutarse

	duration := time.Since(start).Milliseconds()
	fmt.Printf("%d ms\n", duration)

} //La aplicación va a terminar hasta que todas las gouroutines hayan finalizado

//https://www.youtube.com/watch?v=hnBy2kA3MJU
