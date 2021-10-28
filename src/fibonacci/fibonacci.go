package fibonacci

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Main() {
	c := make(chan int)
	quit := make(chan int)

	// go func()... se lanza de manera concurrente:
	// la línea fmt.Println(<-c) queda en espera hasta que haya algo que leer de <-c.
	// Al ser lanzado concurrentemente, mientras dicha línea espera, se va a la ejecución de fibonacci(c, quit),
	// que guardará valores en c <- x tantas veces se produzca la espera de la línea fmt.Println(<-c).
	// Una vez sale del bucle, pasa a esperar case <-quit, que será satisfecho con quit <- 0, por lo que imprimirá "quit" y finalizará
	// el programa.
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
