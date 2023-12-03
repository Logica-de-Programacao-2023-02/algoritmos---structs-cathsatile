package main

import (
	"fmt"
	"time"
)

type Viagem struct {
	Origem  string
	Destino string
	Data    time.Time
	Preco   float64
}

func encontrarViagemMaisCara(viagens []Viagem) (Viagem, error) {
	if len(viagens) == 0 {
		return Viagem{}, fmt.Errorf("O slice de viagens está vazio")
	}

	viagemMaisCara := viagens[0]

	for _, viagem := range viagens[1:] {
		if viagem.Preco > viagemMaisCara.Preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara, nil
}

func main() {
	viagensExemplo := []Viagem{
		{Origem: "Cidade A", Destino: "Cidade B", Data: time.Now(), Preco: 500.0},
		{Origem: "Cidade C", Destino: "Cidade D", Data: time.Now().AddDate(0, 1, 0), Preco: 750.0},
		{Origem: "Cidade E", Destino: "Cidade F", Data: time.Now().AddDate(0, 2, 0), Preco: 1000.0},
		{Origem: "Cidade G", Destino: "Cidade H", Data: time.Now().AddDate(0, 3, 0), Preco: 800.0},
	}

	viagemMaisCara, err := encontrarViagemMaisCara(viagensExemplo)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Printf("Viagem mais cara:\n")
		fmt.Printf("Origem: %s\n", viagemMaisCara.Origem)
		fmt.Printf("Destino: %s\n", viagemMaisCara.Destino)
		fmt.Printf("Data: %s\n", viagemMaisCara.Data.Format("02/01/2006"))
		fmt.Printf("Preço: R$ %.2f\n", viagemMaisCara.Preco)
	}
}
