package main

import (
	"fmt"
	"strings"
)

// Opción 1: Función para calcular el promedio de las notas
func averageGrade(notas []float64) float64 {
	if len(notas) == 0 {
		return 0
	}
	var suma float64
	for _, nota := range notas {
		suma += nota
	}
	return suma / float64(len(notas))
}

// Opción 1: Procesamiento completo del curso
func procesarCurso() {
	var cantidad int
	fmt.Print("\nIngrese la cantidad de estudiantes del curso: ")
	fmt.Scan(&cantidad)

	if cantidad <= 0 {
		fmt.Println("La cantidad de estudiantes debe ser mayor a 0.")
		return
	}

	notas := make([]float64, cantidad)

	for i := 0; i < cantidad; i++ {
		for {
			var nota float64
			fmt.Printf("Ingrese la nota del estudiante %d (0 a 100): ", i+1)
			fmt.Scan(&nota)

			if nota >= 0 && nota <= 100 {
				notas[i] = nota
				break
			} else {
				fmt.Println("Error: La nota debe estar entre 0 y 100.")
			}
		}
	}

	promedio := averageGrade(notas)
	fmt.Printf("\n--- RESULTADOS DEL CURSO ---\n")
	fmt.Printf("Promedio general: %.2f\n", promedio)

	// Determinación de estado con IF
	if promedio >= 70 {
		fmt.Println("Estado: APROBADO")
	} else {
		fmt.Println("Estado: REPROBADO")
	}

	// Evaluación de desempeño con SWITCH
	switch {
	case promedio >= 90 && promedio <= 100:
		fmt.Println("Desempeño: Excellent performance")
	case promedio >= 80 && promedio < 90:
		fmt.Println("Desempeño: Good performance")
	case promedio >= 70 && promedio < 80:
		fmt.Println("Desempeño: Satisfactory performance")
	case promedio < 70:
		fmt.Println("Desempeño: Needs improvement")
	}
}

// Opción 2: Función para calcular la suma de 1 a n
func sumarHastaN() {
	var n int
	fmt.Print("\nIngrese un número entero positivo (n): ")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Por favor ingrese un número mayor o igual a 1.")
		return
	}

	suma := 0
	for i := 1; i <= n; i++ {
		suma += i
	}

	fmt.Printf("La suma de los números del 1 al %d es: %d\n", n, suma)
}

func main() {
	var opcion string

	for {
		fmt.Println("\n==================================")
		fmt.Println("         MENÚ PRINCIPAL           ")
		fmt.Println("==================================")
		fmt.Println("1. Calcular promedio del curso")
		fmt.Println("2. Sumar números de 1 a N")
		fmt.Println("0 o salir. Salir del programa")
		fmt.Print("Seleccione una opción: ")

		fmt.Scan(&opcion)
		opcion = strings.ToLower(opcion)

		if opcion == "0" || opcion == "salir" {
			fmt.Println("Saliendo del programa... ¡Hasta luego!")
			break
		}

		switch opcion {
		case "1":
			procesarCurso()
		case "2":
			sumarHastaN()
		default:
			fmt.Println("Opción no válida. Intente nuevamente.")
		}
	}
}