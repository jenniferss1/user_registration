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
