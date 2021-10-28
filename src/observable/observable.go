package observable

import (
	"fmt"
	"time"
)

func emisorA(output chan<- string) {
	for {
		output <- "Emisor: A"
		time.Sleep(time.Second * 1)
	}
}

func emisorB(output chan<- string) {
	for {
		output <- "Emisor: b"
		time.Sleep(time.Second * 3)
	}
}

func Main() {
	canalA := make(chan string)
	canalB := make(chan string)

	go emisorA(canalA)
	go emisorB(canalB)

	for {
		select {
		case cA := <-canalA:
			fmt.Println(cA)
		case cB := <-canalB:
			fmt.Println(cB)
		}
	}

}
