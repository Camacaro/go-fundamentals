package main

import "fmt"

// Con esto definimos una funcion que recibe un string y un channel de tipo string
// y ademas el channel solo recibira datos sera un canal de entradas de datos
// (Para decir que es un channel de salida se hace con c <-chan string)
func say(text string, c chan<- string) {
	// Con esta forma lo que hacemos es que vamos a ingresar un dato al channel
	c <- text

	// Para sacar el dato de un channel y asignar a una variable se hace de la siguiente forma
	// var variable string = <-c
	// text = <-c
}

func printChannelOutput(c <-chan string) { // canal de salida de datos
	var output string
	output = <-c
	fmt.Println(output)
}

// Las waitgroup son mucho mas eficientes que los channels
// si tu codigo no necesita de una eficiencia optima se pueden usar channels
// Los channels son una forma de comunicacion entre procesos de las goroutines
// ya que ellos manejan de forma nativa la comunicacion entre ellos y ademas de datos primitivos como los waitgroups
// para manejar las concurrencias.
func main() {

	// c := make(chan string) // dinamically accepts goroutines
	// Se crea un chanel de tipo string y va recibir de manera simultanea una(1) goroutine
	c := make(chan string, 1)

	fmt.Println("Hello")

	go say("Bye", c)

	// De esta forma lo que hacemos es obtener el dato que se ingreso en el channel
	fmt.Println(<-c)

	// printChannelOutput(c)
}

// Goroutines sería útil a la hora de procesar documentos o imágenes.
// La concurrencia en go la utilizaría para manejar múltiples solicitudes a un servicio o servidor.
