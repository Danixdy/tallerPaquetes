// main.go: Menú interactivo para probar el paquete 'funciones'.
// Adaptado de lo anterior; usa ConversorMonedas e ImprimirConteoVocales.
//
// Para ejecutar localmente: go run .
//
// Documentación completa en README.md y comentarios en funciones/*.go.
// Repositorio: https://github.com/Danixdy/tallerPaquetes
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Danixdy/tallerPaquetes/funciones" // Import del paquete adaptado
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("=== Taller Práctico: Paquete de Utilidades en Go ===")
	fmt.Println("Bienvenido. Elija una opción:")
	fmt.Println("1. Conversor de Monedas")
	fmt.Println("2. Contador de Vocales")
	fmt.Println("3. Salir")

	for {
		fmt.Print("\nOpción (1-3): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		switch input {
		case "1":
			funciones.ConversorMonedas() // Interactivo, como en lo anterior
		case "2":
			fmt.Print("Ingrese una frase: ")
			scanner.Scan()
			frase := scanner.Text()
			funciones.ImprimirConteoVocales(frase) // Imprime conteo
		case "3":
			fmt.Println("¡Hasta luego!")
			return
		default:
			fmt.Println("Opción inválida. Intente de nuevo.")
		}
	}
}
