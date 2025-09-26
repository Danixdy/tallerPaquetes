package funciones

import (
	"fmt"
	"strings"
)

// ConversorMonedas prompts the user for a USD amount and converts it...
func ConversorMonedas() {
	var dolares float64
	var opcion string

	fmt.Print("Ingrese el valor en dólares a convertir: ")
	fmt.Scan(&dolares)

	fmt.Println(`Elija la moneda:
    1. Euro
    2. LB (Libras)
    3. Won
    4. BTC
    Ingrese el número: `)
	fmt.Scan(&opcion)

	switch opcion {
	case "1":
		euros := dolares * 0.85
		fmt.Printf("%.2f USD = %.2f EUR\n", dolares, euros)
	case "2":
		libras := dolares * 0.74
		fmt.Printf("%.2f USD = %.2f GBP\n", dolares, libras)
	case "3":
		won := dolares * 1370.50
		fmt.Printf("%.2f USD = %.2f KRW\n", dolares, won)
	case "4":
		btc := dolares * 0.000016
		fmt.Printf("%.2f USD = %.6f BTC\n", dolares, btc)
	default:
		fmt.Println("Opción inválida.")
	}
}

// contarCadaVocal counts vowels in text (case-insensitive).
func ContarCadaVocal(texto string) map[rune]int {
	resultado := map[rune]int{'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0}
	textoLower := strings.ToLower(texto)
	for _, c := range textoLower {
		if count, ok := resultado[c]; ok {
			resultado[c] = count + 1
		}
	}
	return resultado
}
func ImprimirConteoVocales(texto string) {
	conteo := ContarCadaVocal(texto)
	fmt.Printf("Conteo de vocales en '%s':\n", texto)
	for vocal, cantidad := range conteo {
		if cantidad > 0 {
			fmt.Printf("  %c: %d\n", vocal, cantidad)
		}
	}

}
