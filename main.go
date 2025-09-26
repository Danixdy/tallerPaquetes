package main

import (
	"fmt"

	"github.com/Danixdy/tallerPaquetes/funciones" // Importa tu paquete (ajusta 'tuusuario')
)

func main() {
	fmt.Println("=== Bienvenido al Paquete Resolucion (desde tallerPaquetes) ===")
	fmt.Println("Este programa importa y usa las funciones del paquete 'resolucion'.")
	// Prueba 1: Conversor de Monedas
	// Nota: La función es interactiva, pedirá input al usuario.
	fmt.Println("\n--- 1. Probando Conversor de Monedas ---")
	funciones.ConversorMonedas() // Llama a la función de tu paquete
	// Prueba 2: Contador de Vocales
	// Usa una frase de ejemplo y muestra el conteo.
	fmt.Println("\n--- 2. Probando Contador de Vocales ---")
	frase := "¡Hola, mundo! Esto es una prueba con vocales en mayúsculas y minúsculas: AeIoU."

	// Opción A: Usa directamente la función que devuelve el map (tu código original)
	conteo := funciones.contarCadaVocal(frase) // Ajusta el nombre si lo cambiaste a ContarVocales
	fmt.Printf("Frase analizada: '%s'\n", frase)
	fmt.Println("Conteo de vocales:")
	for vocal, cantidad := range conteo {
		fmt.Printf("  %c: %d veces\n", vocal, cantidad)
	}
	// Opción B: Si agregaste ImprimirConteoVocales en el código corregido, úsala:
	// resolucion.ImprimirConteoVocales(frase)
	fmt.Println("\n=== Fin de la demostración ===")
}
