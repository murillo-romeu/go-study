/*
Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/
package main

var rest int

func main() {
	for i := 10; i <= 100; i++ {
		rest = i % 4
		println(i, " / 4 resto = ", rest)
	}
}
