package di

import (
	repository "github.com/miyamoto2024/cli-pokemon-api/domain/repository"
	"github.com/miyamoto2024/cli-pokemon-api/infrastructure"
	usecase "github.com/miyamoto2024/cli-pokemon-api/usecase"
	"gorm.io/gorm"
)

type CRUDRegistory interface {
	PokemonRepo() repository.PokemonRepository
	PokemonCRUDUsecase() usecase.PokemonDataCRUDUsecaseInterface
}
type crudRegistory struct {
	db *gorm.DB
}

func NewCRUDRegistory(db *gorm.DB) CRUDRegistory {
	return &crudRegistory{db: db}
}
func (r *crudRegistory) PokemonRepo() repository.PokemonRepository {
	return infrastructure.NewPokemonRepository(r.db)
}
func (r *crudRegistory) PokemonCRUDUsecase() usecase.PokemonDataCRUDUsecaseInterface {
	return usecase.NewPokemonDataCRUDUsecase(r.PokemonRepo())
}

type FetchRegistory interface {
	PokemonAPIRepo() repository.PokemonAPIRepository
	PokemonFetchUsecase() usecase.PokemonDataFetachUsecaseInterface
}
type fetchRegistory struct {
	id int
}

func NewFetchRegistory(id int) FetchRegistory {
	return &fetchRegistory{id: id}
}

func (r *fetchRegistory) PokemonAPIRepo() repository.PokemonAPIRepository {
	return infrastructure.NewPokemonAPI(r.id)
}
func (r *fetchRegistory) PokemonFetchUsecase() usecase.PokemonDataFetachUsecaseInterface {
	return usecase.NewPokemonDataFetchUsecase(r.PokemonAPIRepo())
}
