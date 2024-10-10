package infrastructure_test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/miyamoto2024/cli-pokemon-api/domain/model"
	"github.com/miyamoto2024/cli-pokemon-api/domain/repository"
	"github.com/miyamoto2024/cli-pokemon-api/infrastructure"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupMockDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock database: %v", err)
	}
	dialector := mysql.New(mysql.Config{
		Conn:                      mockDB,
		SkipInitializeWithVersion: true,
	})
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm database: %v", err)
	}
	return db, mock
}

// NewPokemonRepository関数がrepository.PokemonRepositoryインターフェースを実装しているか
func TestNewPokemonRepository(t *testing.T) {
	db, _ := setupMockDB(t)
	repo := infrastructure.NewPokemonRepository(db)
	assert.Implements(t, (*repository.PokemonRepository)(nil), repo)
}

// PokemonRepositoryのGetAllメソッド（SELECT）が期待通りに動作するか
func TestPokemonRepository_GetAll(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := infrastructure.NewPokemonRepository(db)
	// テストデータをモック設定
	rows := sqlmock.NewRows([]string{"id", "name", "height", "weight", "base_experience"}).
		AddRow(1, "bulbasaur", 7, 69, 64).
		AddRow(2, "ivysaur", 10, 130, 142)
	// 期待しているSELECT文が発行されたらモックからデータを返す
	mock.ExpectQuery("SELECT \\* FROM `pokemons`").WillReturnRows(rows)
	pokemons, err := repo.GetAll()
	assert.NoError(t, err)
	// mockデータと取得データが全て一致しているか
	assert.Equal(t, pokemons, []model.Pokemon{
		{ID: 1, Name: "bulbasaur", Height: 7, Weight: 69, BaseExperience: 64},
		{ID: 2, Name: "ivysaur", Height: 10, Weight: 130, BaseExperience: 142},
	})
}

// PokemonRepositoryのCreateメソッド（INSERT）が期待通りに動作するか
func TestPokemonRepository_Create(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := infrastructure.NewPokemonRepository(db)
	// テストデータ
	pokemon := model.Pokemon{ID: 1, Name: "bulbasaur", Height: 7, Weight: 69, BaseExperience: 64}
	// 期待しているINSERT文が発行されたら成功
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `pokemons`").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err := repo.Create(pokemon)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

// PokemonRepositoryのDeleteメソッド（DELETE）が期待通りに動作するか
func TestPokemonRepository_Delete(t *testing.T) {
	db, mock := setupMockDB(t)
	repo := infrastructure.NewPokemonRepository(db)
	sqlmock.NewRows([]string{"id", "name", "height", "weight", "base_experience"}).
		AddRow(1, "bulbasaur", 7, 69, 64).
		AddRow(2, "ivysaur", 10, 130, 142)
	// 期待しているDELETE文が発行されたら成功
	mock.ExpectBegin()
	mock.ExpectExec("DELETE FROM `pokemons` WHERE `pokemons`.`id` = ?").
		WithArgs(1).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectCommit()
	err := repo.Delete(1)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}
