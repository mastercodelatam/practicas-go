package main

import "fmt"

func returnInt() int {
	return 10
}

func doubleReturn() (int, int) {	
	return 50, 57
}
func main() {
	println("Hola Mundo")
	//declaracion de contantes
	value1, _ := returnInt(), returnInt()
	fmt.Println(value1)
	value2, value3 := doubleReturn()
	fmt.Println(value2,value3)
}