package repository

import (
	"clean-API/internal/model"
	"database/sql"
	"fmt"
	"log"
)

//go:generate mockgen -source=user.go -destination=user_mock.go -package repository

type UserRepository interface {
	Save(model.User) (model.User, error)
	GetUserByEmail(string) (model.User, error)
	FindByID(int64) (model.User, error)
}

type user struct {
	db *sql.DB
}

func newUserRepository(db *sql.DB) UserRepository {
	return &user{
		db: db,
	}
}

func (u user) Save(user model.User) (model.User, error) {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	stmt, err := u.db.Prepare(query)

	if err != nil {
		return model.User{}, err
	}

	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Println(err)
		}
	}(stmt)

	result, err := stmt.Exec(user.Email, user.Password)

	if err != nil {
		return model.User{}, err
	}

	id, err := result.LastInsertId()
	user.ID = id
	return user, err
}

func (u user) GetUserByEmail(email string) (model.User, error) {
	var user model.User
	query := `SELECT id, password FROM users WHERE email = ?`
	row := u.db.QueryRow(query, email)

	err := row.Scan(&user.ID, &user.Password)

	if err != nil {
		return model.User{}, fmt.Errorf("user with email %s not found", email)
	}

	return user, nil
}

func (u user) FindByID(id int64) (model.User, error) {
	var user model.User
	query := `SELECT id, email, password FROM users WHERE id = ?`
	row := u.db.QueryRow(query, id)

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		return model.User{}, fmt.Errorf("user with id %d not found", id)
	}
	return user, nil
}
