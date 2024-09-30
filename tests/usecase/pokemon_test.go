package usecase_test

import (
	"testing"

	"github.com/miyamoto2024/cli-pokemon-api/domain/model"
	"github.com/miyamoto2024/cli-pokemon-api/usecase"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// PokemonAPIRepositoryのモックを作成
type MockPokemonAPIRepository struct {
	mock.Mock
}

func (m *MockPokemonAPIRepository) Fetch() (model.Pokemon, error) {
	args := m.Called()
	return args.Get(0).(model.Pokemon), args.Error(1)
}

func TestFetchPokemon(t *testing.T) {
	// モックリポジトリを作成
	mockRepo := new(MockPokemonAPIRepository)
	expectedPokemon := model.Pokemon{ID: 1, Name: "Pikachu"}

	// モックリポジトリのFetchメソッドの期待値を設定
	mockRepo.On("Fetch").Return(expectedPokemon, nil)

	// ユースケースを作成
	usecase := usecase.NewPokemonDataFetchUsecase(mockRepo)

	// FetchPokemonメソッドをテスト
	pokemon, err := usecase.FetchPokemon()

	// アサーション
	assert.NoError(t, err)
	assert.Equal(t, expectedPokemon, pokemon)

	// モックの期待値が満たされていることを確認
	mockRepo.AssertExpectations(t)
}
