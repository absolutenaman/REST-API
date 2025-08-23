package models

import (
	"rest-api/db"
	"rest-api/utils"
)

type User struct {
	Id       int64  `json:"id"`
	Email    string `binding:"required" json:"email"`
	Password string `binding:"required" json:"password"`
}

func (u User) AddUser() error {
	addQuery := `INSERT INTO users (email,password) VALUES (?,?)`
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	resultedRow, err := db.DB.Exec(addQuery, u.Email, hashedPassword)
	if err != nil {
		return err
	}
	u.Id, _ = resultedRow.LastInsertId()
	return nil
}
