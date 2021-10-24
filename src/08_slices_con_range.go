// package main

// // import "fmt"
// import (
// 	"fmt"
// 	"strings"
// )

// func isPalindromeCopilot(word string) bool {
// 	for i := 0; i < len(word)/2; i++ {
// 		if word[i] != word[len(word)-1-i] {
// 			fmt.Printf("The word %s is not a palindrome \n", word)
// 			return false
// 		}
// 	}
// 	fmt.Printf("The word %s is a palindrome \n", word)
// 	return true
// }

// func isPalindrome(text string) {
// 	var textReverse string

// 	for i := len(text) - 1; i >= 0; i-- {
// 		textReverse += string(text[i])
// 	}

// 	if strings.ToLower(text) == strings.ToLower(textReverse) {
// 		fmt.Printf("The word %s is a palindrome \n", text)
// 	} else {
// 		fmt.Printf("The word %s is not a palindrome \n", text)
// 	}
// }

// func main() {

// 	fmt.Println("============== Ejercicio 1 - Recorrido de Slices con Range ==================")
// 	slice := []string{"Hola", "que", "tal", "?"}

// 	fmt.Println("============== Inidce, valor")
// 	for indice, value := range slice {
// 		fmt.Println("Inidce, valor", indice, value)
// 	}

// 	fmt.Println("============== Solo valor")
// 	for _, value := range slice {
// 		fmt.Println("valor", value)
// 	}

// 	fmt.Println("============== Solo Indice")
// 	for indice := range slice {
// 		fmt.Println("indice", indice)
// 	}

// 	fmt.Println("============== Ejercicio 2 - isPalindrome ==================")
// 	isPalindrome("radar")
// 	isPalindrome("amor")
// 	isPalindrome("amor a roma")
// 	isPalindrome("casas")
// 	isPalindrome("ama")
// 	isPalindrome("Ama")

// 	fmt.Println("============== Ejercicio 3 - isPalindromeCopilot ==================")
// 	isPalindromeCopilot("radar")
// 	isPalindromeCopilot("amor")
// 	isPalindromeCopilot("amor a roma")
// 	isPalindromeCopilot("casas")
// 	isPalindromeCopilot("ama")
// 	isPalindromeCopilot("Ama")
// }
