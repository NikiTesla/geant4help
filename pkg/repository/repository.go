package repository

import "fmt"

type User struct {
	Name   string  `json:"name"`
	Age    int     `json:"age"`
	Job    string  `json:"job"`
	Salary float64 `json:"salary"`
}

func (u User) String() string {
	return fmt.Sprintf("%s is %d and works in %s for %f",
		u.Name, u.Age, u.Job, u.Salary)
}
