package persistence

import (
	"database/sql"
	"fmt"
	"log"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Repository struct {
	ProductRepository interfaces.ProductRepository
	db                *sql.DB
}

func NewRepository(dbHost, dbPort, dbUser, dbPass, dbName string) (*Repository, error) {
	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost,
		dbPort,
		dbUser,
		dbPass,
		dbName,
	)
	log.Println(conn)
	db, err := sql.Open("pgx", conn)
	if err != nil {
		log.Println("DB error : db.go : ", err)
		return nil, err
	}
	// defer db.Close()
	return &Repository{
		ProductRepository: NewProductRepository(db),
		db:                db,
	}, nil
}

func (repo *Repository) Close() error {
	return repo.db.Close()
}
