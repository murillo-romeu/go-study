/*
Crie um loop utilizando a sintaxe: for {}
Utilize-o para demonstrar os anos desde que você nasceu.
*/
package main

func main() {
	ano := 1986

	for {
		println(ano)
		ano++
		if ano > 2023 {
			break
		}

	}
}
