package main

import "fmt"

type Musica struct {
	Titulo  string
	Artista string
	Duracao float64
}

type Playlist struct {
	Nome    string
	Musicas []Musica
}

func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.Musicas {
			if musica.Titulo == titulo {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}

func main() {

	playlistsExemplo := []Playlist{
		{
			Nome: "Playlist 1",
			Musicas: []Musica{
				{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3.5},
				{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4.2},
			},
		},
		{
			Nome: "Playlist 2",
			Musicas: []Musica{
				{Titulo: "Música 2", Artista: "Artista 2", Duracao: 4.2},
				{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2.8},
			},
		},
		{
			Nome: "Playlist 3",
			Musicas: []Musica{
				{Titulo: "Música 1", Artista: "Artista 1", Duracao: 3.5},
				{Titulo: "Música 3", Artista: "Artista 3", Duracao: 2.8},
			},
		},
	}

	tituloBuscado := "Música 2"
	playlistsEncontradas := buscarPlaylistsPorTitulo(tituloBuscado, playlistsExemplo)

	fmt.Printf("Playlists que contêm a música \"%s\":\n", tituloBuscado)
	for _, playlist := range playlistsEncontradas {
		fmt.Printf("- %s\n", playlist.Nome)
	}
}
