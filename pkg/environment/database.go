package environment

import "database/sql"

type DataBase struct {
	DB *sql.DB
}

func NewDataBase(cfg DBConfig) (*DataBase, error) {
	return &DataBase{}, nil
}
