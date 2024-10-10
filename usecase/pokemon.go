package usecase

import (
	"github.com/miyamoto2024/cli-pokemon-api/domain/model"
	"github.com/miyamoto2024/cli-pokemon-api/domain/repository"
)

type PokemonDataOperationUsecaseInterface interface {
	GetAllPokemons() ([]model.Pokemon, error)
	CreatePokemon(pokemon model.Pokemon) error
	DeletePokemon(id int) error
}
type pokemonDataCRUDUsecase struct {
	repo repository.PokemonRepository
}

func NewPokemonDataOperationUsecase(repo repository.PokemonRepository) PokemonDataOperationUsecaseInterface {
	return &pokemonDataCRUDUsecase{repo: repo}
}
func (pu *pokemonDataCRUDUsecase) GetAllPokemons() ([]model.Pokemon, error) {
	return pu.repo.GetAll()
}
func (pu *pokemonDataCRUDUsecase) CreatePokemon(pokemon model.Pokemon) error {
	return pu.repo.Create(pokemon)
}
func (pu *pokemonDataCRUDUsecase) DeletePokemon(id int) error {
	return pu.repo.Delete(id)
}

// 外部APIを利用する場合
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
