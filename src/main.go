package main

import (
	"fmt"
	"sync"
	"time"
)

func say(txt string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(txt)
}

// Main es una Gorutine
func main() {
	// Esto lo que hace es crear un grupo de gorutinas, y esperar a que todas terminen
	// la variable wg es un puntero a un WaitGroup
	// y la va liberando poco a poco
	var wg sync.WaitGroup

	fmt.Println("Hello")

	// Agregamos una gorutina a la lista de espera
	wg.Add(1)
	go say("World", &wg)

	// Esperamos a que terminen las gorutinas, antes de que muera el main
	wg.Wait()

	go func(text string) {
		fmt.Println("Funcion anonima: ", text)
	}("Adios")

	// Esto no es una buena práctica. Ya que le agregamos un delay a la ejecución
	time.Sleep(time.Second * 1)
}

// Los WaitGroup son una forma de controlar la ejecución de una gorutina
// y esperar a que todas terminen.
// Si no se terminan, el programa se queda esperando.
// Sin embargo apesar de ser muy buenas prácticas (da optimizacion al codigo), no es recomendable usarlo.
// ya que suelen ser un poco complejo mantenerlos a largo plazo.

/*
	Un WaitGroup espera a que una colección de goroutines termine su ejecución.
	Para esto se una la WaitGroup.Add() ( wg.add(1) en el ejemplo de la clase).
	El número entero indica el número de goroutines que debe esperar para finalizar la ejecución de la goroutine principal.

	Cada vez que una goroutine termina su ejecución, llama el método Done(). Esto hace que el contador del WaitGroup se reduzca.
	Cuando el contador llegue a zero la rutina principal continuará su ejecución.

	La función wait() bloquea la rutina principal hasta que todas las demás rutinas del grupo hayan terminado.
*/

// En su defecto es mejor usar los chaneeles. o canales de comunicación. (ver el siguiente ejemplo)
// Donde se comunica una gorutina con otra.
