package controller

import (
	"github.com/miyamoto2024/cli-pokemon-api/di"
)

func DispPokemonFetchByIdController(id int) {
	// DIはdiに隠蔽
	di := di.NewFetchRegistory(id)
	usecase := di.PokemonFetchUsecase()
	pokemon, _ := usecase.FetchPokemon()
	println(pokemon.Name)
}
