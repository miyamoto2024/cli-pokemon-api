package model

import "errors"

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	BaseExperience int    `json:"base_experience"`
}

func NewPokemon(id int, name string, height int, weight int, baseExperience int) (Pokemon, error) {
	// バリデーション
	if id < 0 {
		return Pokemon{}, errors.New("id must be greater than 0")
	}
	if name == "" {
		return Pokemon{}, errors.New("name must not be empty")
	}
	if height < 0 {
		return Pokemon{}, errors.New("height must be greater than 0")
	}
	if weight < 0 {
		return Pokemon{}, errors.New("weight must be greater than 0")
	}
	if baseExperience < 0 {
		return Pokemon{}, errors.New("baseExperience must be greater than 0")
	}
	// ポケモンの構造体を返却
	return Pokemon{
		ID:             id,
		Name:           name,
		Height:         height,
		Weight:         weight,
		BaseExperience: baseExperience,
	}, nil
}
