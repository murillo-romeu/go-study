package main

import (
	"fmt"
)

const (
	x = iota
	_
	_
	z
)

func main() {
	fmt.Printf("%v \n%T\n", z, z)
}
