package models

import (
	"fmt"
	DB "rest-api/db"
	"time"
)

type Events struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Location    string    `json:"location"`
	DateTime    time.Time `json:"dateTime"`
	UserId      int64     `json:"userId"`
}

var events []Events

func GetAllEvents() []Events {

	return events
}
func (e Events) Sava() {
	insertQuery := `INSERT INTO events (name,description,location,dateTime) VALUES (?,?,?,?)`
	res, err := DB.DB.Exec(insertQuery, e.Name, e.Description, e.Location, e.DateTime, e.UserId)
	if err != nil {
		panic(err)
	}
	fmt.Println("!!!!", res)
	e.ID, _ = res.LastInsertId()
	e.UserId, _ = res.LastInsertId()

}
