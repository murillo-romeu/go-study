/*
Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/
package main

import "fmt"

func main() {
	esporteFavorito := "basquete"

	switch esporteFavorito {
	case "volei":
		fmt.Println("Gosta de jogar Volei")
	case "roquei":
		fmt.Println("Gosta de jogar Roquei")
	case "hipismo":
		fmt.Println("Gosta de jogar Hipismo")
	case "basquete":
		fmt.Println("Gosta de jogar basquete")
	}
}
