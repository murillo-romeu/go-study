/*
Crie constantes tipadas e não-tipadas.
*/

package main

import "fmt"

const (
	x float64 = 10
	z         = 11
	y string  = "amor"
)

func main() {
	fmt.Println(x, z, y)
	fmt.Printf("%T %T %T", x, z, y)
}
