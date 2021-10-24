// package main

// import "fmt"

// func main() {

// 	/**
// 	En Go, los arrays poseen un tamaño fijo y son inmutables,
// 	mientras que en los slices su tamaño es dinámico y los puedes modificar.
// 	*/

// 	fmt.Println("============== Ejercicio 1 - Arrays Inmutables ==================")
// 	// NOMBRE_VARIABLE := [LONGITUD]TIPO_DATO
// 	// var nombre_variable [LONGITUD]tipo_dato
// 	var array [4]int
// 	fmt.Println("inicio array", array)

// 	array[0] = 1
// 	array[1] = 2
// 	fmt.Println("con values array", array)

// 	fmt.Println("Longitud", len(array))
// 	fmt.Println("Capacidad maxima del array", cap(array))

// 	fmt.Println("============== Ejercicio 2 - Slice ==================")
// 	// es como un array, pero no se le indica la cantidad de elementos que tendra
// 	slice := []int{0, 1, 2, 3, 4, 5, 6}
// 	fmt.Println("inicio slice", slice)
// 	fmt.Println("Longitud", len(slice))
// 	fmt.Println("Capacidad maxima del slice", cap(slice))

// 	fmt.Println("============== Ejercicio 3 - Metodos en el Slice ==================")
// 	fmt.Println("Primer elemento", slice[0])
// 	//  slice[INCLUSIVO:EXCLUSIVO]
// 	fmt.Println("Primer elemento hasta el indice 3", slice[:3])
// 	fmt.Println("Elemtos desde el indice 2 hasta 4 (No imprime el elemento 4)", slice[2:4])
// 	fmt.Println("Elemtos desde el indice 4", slice[4:])

// 	fmt.Println("============== Ejercicio 4 - Append  ==================")
// 	// Agregar un elemento
// 	slice = append(slice, 7)
// 	fmt.Println("Se agrego un elemento", slice)

// 	// Agregar varios elementos
// 	newSlice := []int{8, 9, 10}
// 	// spread operator (...) para agregar varios elementos destructurando el slice
// 	slice = append(slice, newSlice...)
// 	fmt.Println("Se agrego varios elemento", slice)
// }
