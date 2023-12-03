package main

import "fmt"

type Music struct {
	Titulo  string
	Artista string
	Duracao float64
}

type Play struct {
	Nome    string
	Musicas []Music
}

func imprimirPlaylist(playlist Play) {
	fmt.Printf("Playlist: %s\n", playlist.Nome)

	tempoTotal := 0.0

	for _, musica := range playlist.Musicas {
		fmt.Printf("Música: %s\n", musica.Titulo)
		fmt.Printf("Artista: %s\n", musica.Artista)
		fmt.Printf("Duração: %.2f minutos\n", musica.Duracao)

		tempoTotal += musica.Duracao

		fmt.Println("-----")
	}

	fmt.Printf("Tempo Total da Playlist: %.2f minutos\n", tempoTotal)
}

func main() {

	playlistExemplo := Play{
		Nome: "Minha Playlist",
		Musicas: []Music{
			{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3.5},
			{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4.2},
			{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2.8},
		},
	}

	imprimirPlaylist(playlistExemplo)
}
