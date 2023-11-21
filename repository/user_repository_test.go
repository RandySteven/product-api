package repository_test

import (
	"testing"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/repository"
	"github.com/DATA-DOG/go-sqlmock"
)

func TestUserByTestDB(t *testing.T) {

}

func TestUserByMock(t *testing.T) {
	t.Run("should return users", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		// query := "^UPDATE products SET deleted_at = NOW\\(\\) WHERE id = \\$1$"
		if err != nil {
			t.Fatalf("error creating sqlmock: %v", err)
		}
		defer db.Close()

		_ = repository.NewUserRepository(db)

		mock.ExpectPrepare("").
			ExpectExec().
			WithArgs(1).
			WillReturnResult(sqlmock.NewResult(0, 1))

		// err = repo.Save(nil)
	})
}
