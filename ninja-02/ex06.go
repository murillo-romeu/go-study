/*
Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
Demonstre estes valores.
*/

package main

const (
	_  = iota
	a1 = iota + 2022
	a2
	a3
	a4
)

func main() {
	println(a1, a2, a3, a4)
}
