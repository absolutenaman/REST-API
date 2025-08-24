package models

import (
	"fmt"
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
func (u User) ValidateUser() error {
	getUserQuery := `SELECT * FROM users WHERE EMAIL=?`
	row, err := db.DB.Query(getUserQuery, u.Email)
	if err != nil {
		fmt.Println("err1", err)
		return err
	}
	var user User
	for row.Next() {
		err = row.Scan(&user.Id, &user.Email, &user.Password)
	}
	u.Id = user.Id
	if err != nil {
		fmt.Println("err2", err)
		return err
	}
	_, err = utils.ValidatePassword([]byte(user.Password), []byte(u.Password))
	if err != nil {
		fmt.Println("err3", err)
		return err
	}
	return nil
}
