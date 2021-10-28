package myclass

import "fmt"

func message(text string, c chan<- string) {
	c <- text
}

func Main() {
	fmt.Println("=============== Ejercicio 1 - Range and Close ===============")
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"

	// len Cuantos elementos, goroutines hay en el canal
	// cap Capacidad maxima que puede almacenar el channel
	fmt.Println("Longitud:", len(c))
	fmt.Println("cantidad:", cap(c))

	// Range y close recorre el canal
	// Se cerrar el canal con close, ya no recibira mas elementos
	close(c)

	// Da error si intentamos meter algun valor porque el canal esta cerrado
	// Tambien da error si el canal esta lleno e intentamos meter un valor
	// c <- "Mensaje3"

	// Es recomendable cerrar los channels cuando ya no se van a usar
	// para evitar que el runtime de go se quede esperando valores en ese channel
	for message := range c {
		fmt.Println(message)
	}

	fmt.Println("=============== Ejercicio 2 - Select ===============")

	// Cuando empezamos a manejar multiples canales y no tenemos certeza cual
	// de los channels va a responder primero, podemos usar select
	// select es una manera de manejar multiples canales de manera asincrona

	email1 := make(chan string)
	email2 := make(chan string)
	// Aquí no sabemos cual de los dos canales va a responder primero
	go message("Mensaje1", email1)
	go message("Mensaje2", email2)

	// Hay que tener presente cuantos channels y elementos tendras porque se menejara en un for
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-email1:
			fmt.Println("Email recibido de email 1", msg1)
		case msg2 := <-email2:
			fmt.Println("Email recibido de email 1", msg2)
		}
	}
}
