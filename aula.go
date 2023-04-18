package main

import (
	"fmt"
)

var x interface{}

func main() {

	x = 20.0

	switch x.(type) {
	case int:
		fmt.Println("é int")
	case bool:
		fmt.Println("é bool")
	case string:
		fmt.Println("é string")
	case float64, float32:
		fmt.Println("é float")
	default:
		fmt.Println("não é nenhum")
	}
}
