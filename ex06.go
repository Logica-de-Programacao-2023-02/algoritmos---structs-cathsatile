package main

import (
	"fmt"
)

type Funcionario struct {
	Nome    string
	Salario float64
	Idade   int
}

func (f *Funcionario) aumentarSalario(porcentagem float64) {
	f.Salario *= (1 + porcentagem/100)
}

func (f *Funcionario) diminuirSalario(porcentagem float64) {
	f.Salario *= (1 - porcentagem/100)
}

func (f *Funcionario) tempoDeServico() int {
	idadeInicioTrabalho := 18
	idadeAtual := f.Idade
	return idadeAtual - idadeInicioTrabalho
}

func main() {

	funcionarioExemplo := Funcionario{
		Nome:    "João",
		Salario: 5000.0,
		Idade:   30,
	}

	fmt.Printf("Nome: %s\n", funcionarioExemplo.Nome)
	fmt.Printf("Salário: R$ %.2f\n", funcionarioExemplo.Salario)
	fmt.Printf("Idade: %d anos\n", funcionarioExemplo.Idade)

	funcionarioExemplo.aumentarSalario(10)
	fmt.Printf("Salário após aumento: R$ %.2f\n", funcionarioExemplo.Salario)

	funcionarioExemplo.diminuirSalario(5)
	fmt.Printf("Salário após diminuição: R$ %.2f\n", funcionarioExemplo.Salario)

	tempoServico := funcionarioExemplo.tempoDeServico()
	fmt.Printf("Tempo de serviço: %d anos\n", tempoServico)
}
