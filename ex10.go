package main

import (
	"fmt"
)

type Filme struct {
	Titulo     string
	Diretor    string
	Ano        int
	Avaliacoes []int
}

func (f *Filme) adicionarAvaliacao(avaliacao int) {
	f.Avaliacoes = append(f.Avaliacoes, avaliacao)
}

func (f *Filme) removerAvaliacao(indice int) error {
	if indice < 0 || indice >= len(f.Avaliacoes) {
		return fmt.Errorf("Índice inválido para remoção de avaliação")
	}

	f.Avaliacoes = append(f.Avaliacoes[:indice], f.Avaliacoes[indice+1:]...)
	return nil
}

func (f Filme) calcularMediaAvaliacoes() float64 {
	if len(f.Avaliacoes) == 0 {
		return 0.0
	}

	soma := 0
	for _, avaliacao := range f.Avaliacoes {
		soma += avaliacao
	}

	return float64(soma) / float64(len(f.Avaliacoes))
}

func (f Filme) imprimirInformacoes() {
	fmt.Printf("Título: %s\n", f.Titulo)
	fmt.Printf("Diretor: %s\n", f.Diretor)
	fmt.Printf("Ano: %d\n", f.Ano)
	fmt.Printf("Avaliações: %v\n", f.Avaliacoes)
	fmt.Printf("Média de Avaliações: %.2f\n", f.calcularMediaAvaliacoes())
}

func main() {

	filmeExemplo := Filme{
		Titulo:     "O Poderoso Chefão",
		Diretor:    "Francis Ford Coppola",
		Ano:        1972,
		Avaliacoes: []int{5, 4, 5, 5, 3},
	}

	fmt.Println("Informações iniciais do filme:")
	filmeExemplo.imprimirInformacoes()

	novaAvaliacao := 4
	filmeExemplo.adicionarAvaliacao(novaAvaliacao)
	fmt.Printf("\nNova avaliação adicionada: %d\n", novaAvaliacao)

	fmt.Println("\nInformações atualizadas do filme:")
	filmeExemplo.imprimirInformacoes()

	indiceAvaliacaoRemovida := 2
	err := filmeExemplo.removerAvaliacao(indiceAvaliacaoRemovida)
	if err != nil {
		fmt.Println("\nErro ao remover avaliação:", err)
	} else {
		fmt.Printf("\nAvaliação removida pelo índice %d\n", indiceAvaliacaoRemovida)
	}

	fmt.Println("\nInformações finais do filme:")
	filmeExemplo.imprimirInformacoes()
}
