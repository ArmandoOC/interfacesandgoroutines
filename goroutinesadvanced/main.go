package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/ArmandoOC/interfacesandgoroutines/goroutinesadvanced/data"
)

func main() {
	start := time.Now()

	wg := &sync.WaitGroup{}
	m := &sync.Mutex{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readBook(i, wg, m)
	}

	wg.Wait() //Hace que se espere hasta que todas las gouritines asociadas al waitgroup terminen de ejecutarse

	duration := time.Since(start).Milliseconds()
	fmt.Printf("%d ms\n", duration)

} //La aplicación va a terminar hasta que todas las gouroutines hayan finalizado

//https://www.youtube.com/watch?v=hnBy2kA3MJU

func readBook(id int, wg *sync.WaitGroup, m *sync.Mutex) {
	data.FinishBook(id, m)

	delay := rand.Intn(800)
	time.Sleep(time.Millisecond * time.Duration(delay))

	wg.Done()
}

//go run main.go
//go run --race main.go  //para averiguar si cae en race condition. To see if there are data races.
//cuando una goruotine está trantando de acceder a un valor que otra goroutine está modificando al mismo tiempo
