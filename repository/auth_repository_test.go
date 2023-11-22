package repository_test

import (
	"log"
	"os"
	"testing"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/configs"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepositoryByTestDB(t *testing.T) {
	t.Run("should return user after user do register", func(t *testing.T) {
		if err := godotenv.Load(); err != nil {
			log.Println("no env got")
		}
		testConfig := models.Config{
			DbHost: os.Getenv("TEST_DB_HOST"),
			DbName: os.Getenv("TEST_DB_NAME"),
			DbPort: os.Getenv("TEST_DB_PORT"),
			DbUser: os.Getenv("TEST_DB_USER"),
			DbPass: os.Getenv("TEST_DB_PASS"),
		}
		user := &models.User{
			Name:     "Test User",
			Email:    "test.user@shopee.com",
			Password: "test_1234",
			RoleID:   1,
		}
		service, _ := configs.NewRepository(&testConfig)

		registeredUser, _ := service.AuthRepository.RegisterUser(user)

		assert.Equal(t, user.Name, registeredUser.Name)
		assert.Equal(t, user.Email, registeredUser.Email)
	})
}

func TestAuthRepositoryByMock(t *testing.T) {

}
