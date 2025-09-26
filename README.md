# tallerPaquetes
# Taller Práctico: Paquete de Utilidades en Go

Este es un paquete modular en Golang para el taller de creación de paquetes. Contiene dos funcionalidades principales:

- **Conversor de Monedas**: Convierte montos en USD a EUR (Euro), GBP (Libras Esterlinas), KRW (Won Surcoreano) o BTC (Bitcoin). Usa tasas fijas aproximadas (octubre 2024). Es interactiva: pide input al usuario vía `fmt.Scan` y usa un `switch` interno para procesar la opción.
- **Contador de Vocales**: Cuenta las ocurrencias de vocales (a, e, i, o, u) en una frase, ignorando mayúsculas/minúsculas (case-insensitive). Retorna un `map[rune]int` para uso programático; la impresión se maneja en `main.go` o se puede extender.

**Repositorio**: https://github.com/Danixdy/tallerPaquetes  
**Autor**: Danixdy  
**Versión Go**: 1.21+ (verifica con `go version`)  
**Fecha de actualización**: Septiembre 2025.

El paquete está diseñado para ser probado localmente con un menú interactivo en `main.go` (usando `switch` para opciones). Las funciones están exportadas con mayúscula (ej: `ConversorMonedas`, `ContarCadaVocal`) para visibilidad desde fuera del paquete, resolviendo issues como "undefined".

## Estructura del Paquete
- `go.mod`: Inicializa el módulo `github.com/Danixdy/tallerPaquetes`.
- `main.go`: Menú interactivo para probar las funciones (usa `bufio.Scanner` y `switch` para navegación).
- `funciones/funciones.go`: Archivo principal del paquete con las funciones (`ConversorMonedas` y `ContarCadaVocal`). Incluye comentarios de documentación en estilo Go-doc (para `go doc`).
  - Opcional: Divide en `currency.go` (conversor) y `vowels.go` (vocales) para modularidad, ambos con `package funciones`.

Ejemplo de estructura en la raíz del repositorio:

tallerPaquetes/
├── go.mod
├── README.md
├── main.go
└── funciones/
    └── funciones.go  # Contiene el paquete con imports ("fmt", "strings") y funciones exportadas

## Comandos 
go mod tidy  # Actualiza dependencias
go run main.go  # Corre y prueba (interactivo para conversor)
git add .  # Agrega README.md, main.go, funciones/funciones.go y todo lo nuevo/modificado
git status  # Verifica: ahora todo estará en verde (staged)
git commit -m "Agrego README.md completo con documentación y comandos Git. Actualizo main.go con menú switch."
git push origin main  # Sube a la rama 'main' en GitHub
