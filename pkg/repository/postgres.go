package repository

import (
	"fmt"

	"github.com/NikiTesla/geant4help/pkg/environment"
)

// type Repo struct {
// 	Env environment.Environment
// }

// TODO think about abstraction level, should I create next abstract level
func CreateUser(name, password_hash string, env *environment.Environment) error {
	rows, err := env.DataBase.DB.Query("INSERT INTO users(username, password_hash) VALUES ($1, $2) RETURNING id",
		name, password_hash)
	if err != nil {
		env.Logger.Error(fmt.Sprintf("can't create user, error: %s", err.Error()))
	}
	for rows.Next() {
		var id int
		rows.Scan(&id)
		fmt.Println(id)
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
