package repositories

import (
	"database/sql"
	"golang-beginner-21/practice/models"
)

type UserRepositoryDB struct {
	DB *sql.DB
}

func NewUserRepositoryDB(db *sql.DB) *UserRepositoryDB {
	return &UserRepositoryDB{DB: db}
}

func (repo *UserRepositoryDB) Login(user models.User) (*models.User, error) {
	sqlStatement := `SELECT id, role FROM users WHERE username = $1 AND password = $2`
	err := repo.DB.QueryRow(sqlStatement, user.Username, user.Password).Scan(&user.ID, &user.Role)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}

	user.Password = "" // Hide password from response
	return &user, nil
}

func (repo *UserRepositoryDB) GetByID(id int) (*models.User, error) {
	var user models.User
	sqlStatement := `SELECT id, username, role FROM users WHERE id = $1`
	err := repo.DB.QueryRow(sqlStatement, id).Scan(&user.ID, &user.Username, &user.Role)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		return nil, err
	}
	return &user, nil
}
