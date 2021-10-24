package main

import "fmt"

func main() {
	/**
	La estructura de datos MAP son más eficientes que los Array o Slices,
	ya que usan de forma nativa concurrencia para ejecutar las operaciones.
	*/
	// Diccionario python - Objectos en Javascript
	fmt.Println("============== Ejercicio 1 - Llave valor con Maps  ==================")
	m := make(map[string]int)

	m["Jose"] = 14
	m["Juan"] = 20

	// map[Jose:14 Juan:20]
	fmt.Println(m)

	fmt.Println("============== Ejercicio 2 - Recorrrer el mapa ==================")
	// Estos recorridos son para iterar, pero no saldran de la misma forma
	// porque no se sabe en que orden se imprimen los valores
	// ya que esto se ejecuta de forma recurrente
	for key, value := range m {
		fmt.Println(key, value)
	}

	fmt.Println("============== Ejercicio 3 - Encontrar el valorw ==================")
	value := m["Jose"]
	fmt.Println("value", value)

	// Maria no existe pero como es un valor int te lanzara el zero value
	// Em int el zero value es 0
	value2, ok := m["Maria"]
	fmt.Println("value2 - zero value - not fount key", value2)
	fmt.Println("Existe el valor", ok)

	value3, ok := m["Juan"]
	fmt.Println("value3", value3)
	fmt.Println("Existe el valor", ok)

	fmt.Println("============== Ejercicio 4 - Otra Forma de definir Map ==================")
	temperature := map[string]int{
		"Earth":  15,
		"Mars":   -65,
		"Saturn": -100,
	}

	fmt.Println("temperature", temperature)
	delete(temperature, "Saturn")
	fmt.Println("temperature", temperature)
}
