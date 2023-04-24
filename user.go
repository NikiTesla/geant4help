package geant4help

import "fmt"

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Job      string `json:"job"`
	Age      int    `json:"age"`
	Salary   int    `json:"salary"`
}

func (u User) String() string {
	return fmt.Sprintf("%s is %d and works in %s for %d",
		u.Name, u.Age, u.Job, u.Salary)
}
