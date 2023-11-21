package repository

import (
	"database/sql"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type AuthRepository struct {
	db *sql.DB
}

// RegisterUser implements interfaces.AuthRepository.
func (repo *AuthRepository) RegisterUser(user *models.User) (*models.User, error) {
	query := "INSERT INTO users (name, email, password, role_id, created_at, updated_at) VALUES " +
		"($1, $2, $3, $4, NOW(), NOW()) RETURNING ID"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var userId uint
	err = stmt.
		QueryRow(&user.Name, &user.Email, &user.Password, &user.RoleID).
		Scan(&userId)
	user.ID = userId
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByEmail implements interfaces.AuthRepository.
func (repo *AuthRepository) LoginUserByEmail(email string) (*models.User, error) {
	query := "SELECT id, name, email, password, role_id FROM users WHERE email = $1"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(email).
		Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Password,
			&user.RoleID,
		)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func NewAuthRepository(db *sql.DB) *AuthRepository {
	return &AuthRepository{db}
}

var _ interfaces.AuthRepository = &AuthRepository{}
