/*
Crie um programa que:

    Atribua um valor int a uma variável
    Demonstre este valor em decimal, binário e hexadecimal
    Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
    Demonstre esta outra variável em decimal, binário e hexadecimal

*/

package main

import "fmt"

func main() {
	num := 22

	fmt.Printf("dec: %d bin: %b hex: %#x \n", num, num, num)

	numx := num << 1

	fmt.Printf("dec: %d bin: %b hex: %#x \n", numx, numx, numx)

}
