package repository

import (
	"database/sql"
	"fmt"
	"modulo/model"
)

type UserRepository struct {
	connection *sql.DB
}

func NewUserRepository(connection *sql.DB) UserRepository {
	return UserRepository{
		connection: connection,
	}
}

func (ur *UserRepository) GetUsers() ([]model.Users, error) {
	query := "SELECT * FROM users"
	rows, err := ur.connection.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Users{}, err
	}

	var userList []model.Users
	var userObj model.Users

	for rows.Next() {
		err = rows.Scan(
			&userObj.ID,
			&userObj.User_name,
			&userObj.User_age,
			&userObj.User_weight,
			&userObj.User_height,
		)

		if err != nil {
			fmt.Println(err)
			return []model.Users{}, err
		}

		userList = append(userList, userObj)
	}

	rows.Close()

	return userList, nil
}

func (ur *UserRepository) CreateUser(user model.Users) (int, error) {
	var id int
	query, err := ur.connection.Prepare("INSERT INTO users" +
		"(user_name, user_age, user_weight, user_height)" +
		" VALUES ($1, $2, $3, $4) RETURNING id")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(user.User_name, user.User_age, user.User_weight, user.User_height).Scan(&id)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	query.Close()
	return id, nil
}

func (ur *UserRepository) GetUserById(userId int) (*model.Users, error) {
	query, err := ur.connection.Prepare("SELECT * FROM users WHERE id = $1")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var user model.Users

	err = query.QueryRow(userId).Scan(
		&user.ID,
		&user.User_name,
		&user.User_age,
		&user.User_weight,
		&user.User_height,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	query.Close()

	return &user, nil
}
