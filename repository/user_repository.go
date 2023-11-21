package repository

import (
	"database/sql"

	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/interfaces"
	"git.garena.com/bootcamp/batch-02/shared-projects/product-api.git/models"
)

type UserRepository struct {
	db *sql.DB
}

// Find implements interfaces.UserRepository.
func (repo *UserRepository) Find() ([]models.User, error) {
	query := "SELECT id, name, email, password, role_id FROM users"
	rows, err := repo.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.RoleID)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

// GetUserByEmailAndPassword implements interfaces.UserRepository.
func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
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

// GetUserById implements interfaces.UserRepository.
func (repo *UserRepository) GetUserById(id uint) (*models.User, error) {
	query := "SELECT id, name, email, password, role_id FROM users WHERE id = $1"
	stmt, err := repo.db.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var user models.User
	err = stmt.QueryRow(query, id).
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

// Save implements interfaces.UserRepository.
func (repo *UserRepository) Save(user *models.User) (*models.User, error) {
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

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db}
}

var _ interfaces.UserRepository = &UserRepository{}
