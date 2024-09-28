package repository

import "github.com/miyamoto2024/cli-pokemon-api/domain/model"

type PokemonRepository interface {
	GetAll() ([]model.Pokemon, error)
	Save(pokemon model.Pokemon) error
	Delete(id int) error
}

type PokemonAPIRepository interface {
	Fetch() (model.Pokemon, error)
}
