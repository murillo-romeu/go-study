/*
Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

package main

import "fmt"

func main() {
	num := 100

	fmt.Printf("%d\n%b\n%#x\n", num, num, num)
}
