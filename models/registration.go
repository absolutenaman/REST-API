package models

import "rest-api/db"

type Registrations struct {
	id      int64 `json:"id"`
	userId  int64 `json:"userId"`
	eventId int64 `json:"eventId"`
}

func Register(userId, eventId int64) error {
	query := `INSERT INTO registrations (userId,eventId)  VALUES (?,?)`
	_, err := db.DB.Exec(query, userId, eventId)
	if err != nil {
		return err
	}
	return nil
}
func Cancellation(userId, eventId int64) error {
	query := `DELETE FROM registrations WHERE userId=? AND eventId=?`
	_, err := db.DB.Exec(query, userId, eventId)
	if err != nil {
		return err
	}
	return nil
}
