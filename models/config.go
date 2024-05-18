package models

import (
	"database/sql"

	"time-keeping.com/lib"
)

type Config struct {
	ID    string `json:"id"`
	Field string `json:"field"`
	Value string `json:"value"`
}

func SetConfig(field, value string) error {
	conn := lib.DBConn
	_, err := conn.Exec("INSERT INTO config (field, value) VALUES ($1, $2)", field, value)
	return err
}

func UpdateConfig(field, value string) error {
	conn := lib.DBConn
	_, err := conn.Exec("UPDATE config SET value = $1 WHERE field = $2", value, field)
	return err
}

func GetConfig(field string) (*Config, error) {
	conn := lib.DBConn
	row := conn.QueryRow("SELECT * FROM config WHERE field = $1", field)

	var a Config
	err := row.Scan(&a.ID, &a.Field, &a.Value)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

func ExistConfig(field string) (bool, error) {
	conn := lib.DBConn
	row := conn.QueryRow("SELECT 1 FROM config WHERE field = $1", field)

	var exists int
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return exists > 0, nil
}
