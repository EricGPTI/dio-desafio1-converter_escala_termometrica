package main

import (
	"fmt"
)
// Converter o ponto de ebulição de Kelvin para Celsius
// Formula C=(k - 273,15)
const ebulicaoK float64 = 373.2

func main () {
	C := ebulicaoK - 273.15
	fmt.Printf("A temperatura de ebulição em Celsius é: %.2f", C)
}

