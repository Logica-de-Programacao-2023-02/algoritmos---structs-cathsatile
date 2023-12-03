package main

import "fmt"

type Endereco struct {
	Rua    string
	Numero int
	Cidade string
	Estado string
}

type Pessoa struct {
	Nome     string
	Idade    int
	Endereco Endereco
}

func imprimirEndereco(pessoa Pessoa) {
	fmt.Printf("Endereço de %s:\n", pessoa.Nome)
	fmt.Printf("Rua: %s, Número: %d\n", pessoa.Endereco.Rua, pessoa.Endereco.Numero)
	fmt.Printf("Cidade: %s, Estado: %s\n", pessoa.Endereco.Cidade, pessoa.Endereco.Estado)
}

func main() {
	pessoaExemplo := Pessoa{
		Nome:  "Alice",
		Idade: 25,
		Endereco: Endereco{
			Rua:    "Rua Principal",
			Numero: 123,
			Cidade: "Cidade Exemplo",
			Estado: "Estado Exemplo",
		},
	}
	imprimirEndereco(pessoaExemplo)
}
