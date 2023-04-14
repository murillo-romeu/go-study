/*
Crie um tipo. O tipo subjacente deve ser int.
Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
Na função main:

    Demonstre o valor da variável "x"
    Demonstre o tipo da variável "x"
    Atribua 42 à variável "x" utilizando o operador "="
    Demonstre o valor da variável "x"


*/

package main

import "fmt"

type newtype int

var x newtype

func main() {

	fmt.Println("Valor:", x)
	fmt.Printf("Tipo: %T\n", x)

	x = 42

	fmt.Println("Novo Valor:", x)

}
