package models

import (
	DB "rest-api/db"
	"time"
)

type Events struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"dateTime"`
}

func GetAllEvents() []Events {
	rows, err := DB.DB.Query(`SELECT * FROM events`)
	if err != nil {
		panic(err)
	}
	var arr []Events
	for rows.Next() {
		var event Events
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime)
		if err != nil {
			panic(err)
		}
		arr = append(arr, event)
	}
	defer rows.Close()
	return arr
}
func (e Events) Sava() {
	insertQuery := `INSERT INTO events (name,description,location,dateTime) VALUES (?,?,?,?)`
	res, err := DB.DB.Exec(insertQuery, e.Name, e.Description, e.Location, e.DateTime)
	if err != nil {
		panic(err)
	}
	e.ID, _ = res.LastInsertId()
}

// in case of post request that will be a struct method
// incase of get request it will be a function
func GetAllEventsById(id int64) (Events, error) {
	selectQuery := `SELECT * FROM events WHERE ID =(?)`
	rows, err := DB.DB.Query(selectQuery, id)
	var e Events
	if err != nil {
		return e, err
	}
	for rows.Next() {

		err := rows.Scan(&e.ID, &e.Name, &e.Description, &e.Location, &e.DateTime)
		if err != nil {
			return e, err
		}
	}
	return e, nil
}
func (e Events) UpdateEvent() error {
	query := `UPDATE events SET name=?,description=?,location=?,dateTime=? WHERE id=?`
	_, err := DB.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.ID)
	if err != nil {
		return err
	}
	return nil
}
func (e Events) DeleteEvent() error {
	query := `DELETE FROM EVENTS WHERE id=?`
	_, err := DB.DB.Exec(query, e.ID)
	if err != nil {
		return err
	}
	return nil
}
