package infrastructure

import (
	"github.com/miyamoto2024/cli-pokemon-api/domain/model"
	"github.com/miyamoto2024/cli-pokemon-api/domain/repository"
	"gorm.io/gorm"
)

type pokemonRepository struct {
	db *gorm.DB
}

func NewPokemonRepository(db *gorm.DB) repository.PokemonRepository {
	return &pokemonRepository{db: db}
}

func (r *pokemonRepository) GetAll() ([]model.Pokemon, error) {
	var pokemons []model.Pokemon
	if err := r.db.Find(&pokemons).Error; err != nil {
		return nil, err
	}
	return pokemons, nil
}

func (r *pokemonRepository) Create(pokemon model.Pokemon) error {
	if err := r.db.Create(&pokemon).Error; err != nil {
		return err
	}
	return nil
}

func (r *pokemonRepository) Delete(id int) error {
	if err := r.db.Delete(&model.Pokemon{}, id).Error; err != nil {
		return err
	}
	return nil
}
