package configs

import (
	"fmt"
	"log"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/repository"
	_ "github.com/jackc/pgx/v4/stdlib"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Repository struct {
	ProductRepository interfaces.ProductRepository
	UserRepository    interfaces.UserRepository
	AuthRepository    interfaces.AuthRepository
	db                *gorm.DB
}

func NewRepository(config *models.Config) (*Repository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.DbHost,
		config.DbPort,
		config.DbUser,
		config.DbPass,
		config.DbName,
	)
	log.Println(conn)
	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Println("DB error : db.go : ", err)
		return nil, err
	}
	// defer db.Close()
	return &Repository{
		ProductRepository: repository.NewProductRepository(db),
		UserRepository:    repository.NewUserRepository(db),
		AuthRepository:    repository.NewAuthRepository(db),
		db:                db,
	}, nil
}

func (repo *Repository) Close() <-chan struct{} {
	return repo.db.Statement.Context.Done()
}
