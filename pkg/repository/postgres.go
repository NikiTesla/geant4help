package repository

import (
	"fmt"

	"github.com/NikiTesla/geant4help"
	"github.com/NikiTesla/geant4help/pkg/environment"
)

// TODO think about abstraction level, should I create next abstract level
func CreateUser(name, password_hash string, env *environment.Environment) error {
	rows, err := env.DataBase.DB.Query("INSERT INTO users(username, password_hash) VALUES ($1, $2) RETURNING id",
		name, password_hash)
	if err != nil {
		env.Logger.Error(fmt.Sprintf("can't create user, error: %s", err.Error()))
		return err
	}

	for rows.Next() {
		var id int
		rows.Scan(&id)

		_, err := env.DataBase.DB.Query("INSERT INTO users_info(id) VALUES ($1)", id)
		if err != nil {
			env.Logger.Error(fmt.Sprintf("can't create user, error: %s", err.Error()))
			return err
		}
	}

	return nil
}

func FindUserByUsername(username string, env *environment.Environment) (int, string, error) {
	rawQuery := env.DataBase.DB.QueryRow("SELECT id, password_hash FROM users WHERE username=$1", username)

	var password_hash string
	var id int
	if err := rawQuery.Scan(&id, &password_hash); err != nil {
		return -1, "", err
	}

	return id, password_hash, nil
}

func FindUserByID(id int, env *environment.Environment) (*geant4help.User, error) {
	query := "SELECT username, name, email, job, age, salary from users join users_info using (id) WHERE id=$1"
	row := env.DataBase.DB.QueryRow(query, id)

	var user geant4help.User
	if err := row.Scan(&user.Username, &user.Name, &user.Email, &user.Job, &user.Age, &user.Salary); err != nil {
		env.Logger.Error(err.Error())
		return nil, err
	}

	return &user, nil
}

func EditUserInfo(id int, name, email string, age int, job string, env *environment.Environment) error {
	query := "UPDATE users_info SET name=$1, email=$2, age=$3, job=$4 WHERE id=$5"

	_, err := env.DataBase.DB.Query(query, name, email, age, job, id)
	if err != nil {
		env.Logger.Error(fmt.Sprintf("can't edit user info, error: %s", err.Error()))
		return err
	}

	return nil
}
