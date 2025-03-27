package main

import "fmt"

func main() {

	//declaracion de contantes
	const pi float64 = 3.14159265359
	const pi2 = 3.14

	fmt.Println(pi)
	fmt.Println(pi2)

	//declarion de variables enteras

	base := 12
	var altura int = 14
	var area int

	fmt.Println(base, altura, area)

	var a int
	var b float64
	var c string
	var d bool

	fmt.Println(a, b, c, d)

	// Area de un cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area cuadrado", areaCuadrado)

	//suma
	x := 10
	y := 50

	result := x + y
	fmt.Println("Suma", result)

	//resta
	result = y - x
	fmt.Println("Resta", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multiplicacion", result)

	//division
	result = y / x
	fmt.Println("Division", result)

	//modulo
	result = y % x
	fmt.Println("Modulo", result)

	//incremental
	x++
	fmt.Println("Incremental", x)

	//decremental
	x--
	fmt.Println("Decremental", x)

	//retos
	//1. Calcular el area de un rectangulo
	//2. Calcular el area de un trapecio
	//3. Calcular el area de un circulo

	//Area de un rectangulo
	baseRectangulo := 20
	alturaRectangulo := 10
	areaRectangulo := baseRectangulo * alturaRectangulo
	fmt.Println("Area rectangulo", areaRectangulo)

	//Area de un trapecio
	baseMenorTrapecio := 10
	baseMayorTrapecio := 20
	alturaTrapecio := 10
	areaTrapecio := ((baseMenorTrapecio + baseMayorTrapecio) * alturaTrapecio) / 2
	fmt.Println("Area trapecio", areaTrapecio)

	//Area de un circulo
	radioCirculo := 10
	areaCirculo := pi * (float64(radioCirculo) * float64(radioCirculo))
	fmt.Println("Area circulo", areaCirculo)

	//println
	nombre1 := "Juan"
	edad1 := 30
	fmt.Println(nombre1, edad1)

	//printf
	nombre := "Juan"
	edad := 30
	fmt.Printf("%s tiene %d años de edad\n", nombre, edad)

}
