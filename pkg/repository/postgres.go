package repository

import (
	"fmt"

	"github.com/NikiTesla/geant4help/pkg/environment"
)

// type Repo struct {
// 	Env environment.Environment
// }

// TODO think about abstraction level, should I create next abstract level
func CreateUser(name string, age int, salary float64, env *environment.Environment) error {
	_, err := env.DataBase.DB.Exec("INSERT INTO users(name, age, salary) VALUES ($1, $2, $3)",
		name, age, salary)

	if err != nil {
		env.Logger.Error(fmt.Sprintf("can't create user, error: %s", err.Error()))
	}

	return nil
}

func FindUserByID(id int, env *environment.Environment) (*User, error) {
	if id < 0 {
		return nil, fmt.Errorf("id can't be less than zero")
	}

	rawQuery := env.DataBase.DB.QueryRow("SELECT name, age, job, salary FROM users WHERE id=$1", id)

	var user User
	if err := rawQuery.Scan(&user.Name, &user.Age, &user.Job, &user.Salary); err != nil {
		return &User{}, err
	}

	return &user, nil
}
