package infrastructure

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/miyamoto2024/cli-pokemon-api/domain/model"
	"github.com/miyamoto2024/cli-pokemon-api/domain/repository"
)

const baseURL = "https://pokeapi.co/api/v2/pokemon/"

type pokemonAPIRepository struct {
	pokemonID int
}

func NewPokemonAPI(pokemonID int) repository.PokemonAPIRepository {
	return &pokemonAPIRepository{pokemonID: pokemonID}
}

func (api *pokemonAPIRepository) Fetch() (model.Pokemon, error) {
	resp, err := http.Get(fmt.Sprintf("%s%d", baseURL, api.pokemonID))
	if err != nil {
		return model.Pokemon{}, err
	}
	defer resp.Body.Close()

	var pokemon model.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return model.Pokemon{}, err
	}

	return pokemon, nil
}

func (api *pokemonAPIRepository) FetchAll() ([]model.Pokemon, error) {
	resp, err := http.Get(fmt.Sprintf("%s%d", baseURL, api.pokemonID))
	if err != nil {
		return []model.Pokemon{}, err
	}
	defer resp.Body.Close()

	var pokemon []model.Pokemon
	if err := json.NewDecoder(resp.Body).Decode(&pokemon); err != nil {
		return []model.Pokemon{}, err
	}

	return pokemon, nil
}
