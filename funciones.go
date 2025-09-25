package funciones

import (
	"fmt"
	"strings"
)

func ConversorMonedas(dolares float64) {

	var opcion string

	fmt.Println("Ingrese el valor a cambiar en dolares: ")
	fmt.Scan(&dolares)

	fmt.Println(`Escoga a que valor quiere cambiar:
	1.- Euro
	2.- LB (Libras Esterlinas)
	3.- Won (Surcoreano)
	4.- BTC
	----Ingresando el numero de la opcion-----`)
	fmt.Scan(&opcion)

	switch opcion {
	case "1":
		euros := dolares * 0.85
		fmt.Printf("La conversion de %.2f dolares a euros es %.2f euros\n", dolares, euros)
	case "2":
		libras := dolares * 0.74
		fmt.Printf("La conversion de %.2f dolares a libras esterlinas es %.2f libras\n", dolares, libras)
	case "3":
		won := dolares * 1370.50 // tasa aproximada
		fmt.Printf("La conversion de %.2f dolares a wones surcoreanos es %.2f wones\n", dolares, won)
	case "4":
		btc := dolares * 0.000016 // tasa aproximada
		fmt.Printf("La conversion de %.2f dolares a Bitcoin es %.6f BTC\n", dolares, btc)
	default:
		fmt.Println("Opción no válida")
	}
}

func contarVocales(texto string) map[rune]int {
	resultado := map[rune]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
	for _, c := range strings.ToLower(texto) { // 4. Convertir a minúsculas
		if _, existe := resultado[c]; existe {
			resultado[c]++
		}
	}
	return resultado
}

func ImprimirConteoVocales(texto string) {
	conteo := contarVocales(texto)
	fmt.Printf("Conteo de vocales en '%s':\n", texto)
	for vocal, count := range conteo {
		fmt.Printf("  %c: %d\n", vocal, count)
	}
}
