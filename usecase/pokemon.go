package usecase

import (
	"github.com/miyamoto2024/cli-pokemon-api/domain/model"
	"github.com/miyamoto2024/cli-pokemon-api/domain/repository"
)

type PokemonDataCRUDUsecaseInterface interface {
	GetAllPokemons() ([]model.Pokemon, error)
	SavePokemon(pokemon model.Pokemon) error
	DeletePokemon(id int) error
}
type pokemonDataCRUDUsecase struct {
	repo repository.PokemonRepository
}

func NewPokemonDataCRUDUsecase(repo repository.PokemonRepository) PokemonDataCRUDUsecaseInterface {
	return &pokemonDataCRUDUsecase{repo: repo}
}
func (pu *pokemonDataCRUDUsecase) GetAllPokemons() ([]model.Pokemon, error) {
	return pu.repo.GetAll()
}
func (pu *pokemonDataCRUDUsecase) SavePokemon(pokemon model.Pokemon) error {
	return pu.repo.Save(pokemon)
}
func (pu *pokemonDataCRUDUsecase) DeletePokemon(id int) error {
	return pu.repo.Delete(id)
}

// 外部APIからデータを取得するusecase
type PokemonDataFetachUsecaseInterface interface {
	FetchPokemon() (model.Pokemon, error)
}
type pokemonDataFetchUsecase struct {
	apirepo repository.PokemonAPIRepository
}

func NewPokemonDataFetchUsecase(apirepo repository.PokemonAPIRepository) PokemonDataFetachUsecaseInterface {
	return &pokemonDataFetchUsecase{apirepo: apirepo}
}
func (pfu *pokemonDataFetchUsecase) FetchPokemon() (model.Pokemon, error) {
	return pfu.apirepo.Fetch()
}
