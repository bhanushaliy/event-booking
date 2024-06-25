package models

import "example.com/rest-api/db"

type Users struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u Users) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(u.Email, u.Password)
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}
