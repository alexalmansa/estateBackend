package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Building struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Longitude float64   `json:"longitude"`
	Latitude  float64   `json:"latitude"`
}

func (i *Building) Create(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("INSERT INTO building (name, address, longitude, latitude, created_at, updated_at) VALUES (?,?,?,?,?,?); ", i.Name, i.Address, i.Longitude, i.Latitude, now, now)
	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating building ")
	}
	return nil
}

func GetBuildings(conn *sql.DB, buildingId string) ([]Building, error) {
	var rows *sql.Rows
	var err error
	if buildingId != "" {
		rows, err = conn.Query("SELECT id, name, address, longitude, latitude, created_at, updated_at FROM building WHERE id = ?", buildingId)

	} else {
		rows, err = conn.Query("SELECT id, name, address, longitude, latitude, created_at, updated_at FROM building ")

	}
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var building []Building
	for rows.Next() {
		i := Building{}
		err = rows.Scan(&i.ID, &i.Name, &i.Address, &i.Longitude, &i.Latitude, &i.CreatedAt, &i.UpdatedAt)
		building = append(building, i)
	}
	return building, nil
}
