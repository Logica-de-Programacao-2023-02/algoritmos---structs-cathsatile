package main

import "fmt"

type Animal struct {
	Nome    string
	Especie string
	Idade   int
	Som     string
}

func (a *Animal) modificarSom(novoSom string) {
	a.Som = novoSom
}

func (a Animal) imprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Espécie: %s\n", a.Especie)
	fmt.Printf("Idade: %d anos\n", a.Idade)
	fmt.Printf("Som que faz: %s\n", a.Som)
}

func main() {

	animalExemplo := Animal{
		Nome:    "Bobby",
		Especie: "Cachorro",
		Idade:   5,
		Som:     "Au Au",
	}

	fmt.Println("Informações iniciais do animal:")
	animalExemplo.imprimirInformacoes()

	novoSom := "Miau"
	animalExemplo.modificarSom(novoSom)
	fmt.Printf("\nSom modificado para: %s\n", novoSom)

	fmt.Println("\nInformações atualizadas do animal:")
	animalExemplo.imprimirInformacoes()
}
