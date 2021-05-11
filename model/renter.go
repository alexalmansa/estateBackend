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
	Nif       string    `json:"nif"`
}

func (i *Renter) Create(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("INSERT INTO renter (name, nif, created_at, updated_at) VALUES (?,?,?,?)", i.Name, i.Nif, now, now)
	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating renter")
	}
	return nil
}
func (i *Renter) UpdateRenter(conn *sql.DB) error {

	fmt.Printf("BUILDING ID: %d ", i.ID)
	now := time.Now()
	row := conn.QueryRow("UPDATE renter SET name = ?, nif = ? , updated_at = ? WHERE id = ?; ", i.Name, i.Nif, now, i.ID)

	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem updating building ")
	}
	return nil
}

func GetRenters(conn *sql.DB, renterId string) ([]Renter, error) {
	var rows *sql.Rows
	var err error
	if renterId != "" {
		rows, err = conn.Query("SELECT id, name, nif, created_at, updated_at FROM renter WHERE id = ?", renterId)

	} else {
		rows, err = conn.Query("SELECT id, name, nif, created_at, updated_at FROM renter ")

	}
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var renter []Renter
	for rows.Next() {
		i := Renter{}
		err = rows.Scan(&i.ID, &i.Name, &i.Nif, &i.CreatedAt, &i.UpdatedAt)
		renter = append(renter, i)
	}
	return renter, nil
}

func DeleteRenter(conn *sql.DB, renterId string) error {
	row := conn.QueryRow("DELETE FROM renter WHERE id = ?", renterId)
	err := row.Scan()
	if err != sql.ErrNoRows {
		fmt.Println(" error deleting items %v", err)
		return err
	}
	return nil
}
