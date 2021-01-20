package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Renter struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

func (i *Renter) Create(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("INSERT INTO renter (name, age, created_at, updated_at) VALUES (?,?,?,?)", i.Name, i.Age, now, now)
	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating renter")
	}
	return nil
}

func GetRenters(conn *sql.DB, renterId string) ([]Renter, error) {
	var rows *sql.Rows
	var err error
	if renterId != "" {
		rows, err = conn.Query("SELECT id, name, age, created_at, updated_at FROM renter WHERE id = ?", renterId)

	} else {
		rows, err = conn.Query("SELECT id, name, age, created_at, updated_at FROM renter ")

	}
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var renter []Renter
	for rows.Next() {
		i := Renter{}
		err = rows.Scan(&i.ID, &i.Name, &i.Age, &i.CreatedAt, &i.UpdatedAt)
		renter = append(renter, i)
	}
	return renter, nil
}
