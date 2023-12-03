package main

import (
	"fmt"
	"math"
)

type Circulo struct {
	Raio float64
}

func calcularArea(circulo Circulo) float64 {
	area := math.Pi * circulo.Raio * circulo.Raio
	return area
}

func main() {
	circuloExemplo := Circulo{Raio: 3.0}

	area := calcularArea(circuloExemplo)

	fmt.Printf("A área do círculo com raio %.2f é %.2f\n", circuloExemplo.Raio, area)
}
