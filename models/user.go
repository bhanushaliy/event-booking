package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

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
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil{
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	return nil
}
