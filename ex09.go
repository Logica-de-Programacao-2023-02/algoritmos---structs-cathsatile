package main

import (
	"fmt"
)

type Aluno struct {
	Nome  string
	Idade int
	Notas []float64
}

func (a *Aluno) adicionarNota(nota float64) {
	a.Notas = append(a.Notas, nota)
}

func (a *Aluno) removerNota(indice int) error {
	if indice < 0 || indice >= len(a.Notas) {
		return fmt.Errorf("Índice inválido para remoção de nota")
	}

	a.Notas = append(a.Notas[:indice], a.Notas[indice+1:]...)
	return nil
}

func (a Aluno) calcularMedia() float64 {
	if len(a.Notas) == 0 {
		return 0.0
	}

	soma := 0.0
	for _, nota := range a.Notas {
		soma += nota
	}

	return soma / float64(len(a.Notas))
}

func (a Aluno) imprimirInformacoes() {
	fmt.Printf("Nome: %s\n", a.Nome)
	fmt.Printf("Idade: %d anos\n", a.Idade)
	fmt.Printf("Média das notas: %.2f\n", a.calcularMedia())
}

func main() {

	alunoExemplo := Aluno{
		Nome:  "Maria",
		Idade: 20,
		Notas: []float64{8.5, 7.0, 9.3},
	}

	fmt.Println("Informações iniciais do aluno:")
	alunoExemplo.imprimirInformacoes()

	novaNota := 8.0
	alunoExemplo.adicionarNota(novaNota)
	fmt.Printf("\nNova nota adicionada: %.2f\n", novaNota)

	fmt.Println("\nInformações atualizadas do aluno:")
	alunoExemplo.imprimirInformacoes()

	indiceNotaRemovida := 1
	err := alunoExemplo.removerNota(indiceNotaRemovida)
	if err != nil {
		fmt.Println("\nErro ao remover nota:", err)
	} else {
		fmt.Printf("\nNota removida pelo índice %d\n", indiceNotaRemovida)
	}

	fmt.Println("\nInformações finais do aluno:")
	alunoExemplo.imprimirInformacoes()
}
