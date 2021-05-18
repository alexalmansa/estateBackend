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
	Longitude float32   `json:"longitude"`
	Latitude  float32   `json:"latitude"`
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
func (i *Building) Update(conn *sql.DB) error {

	fmt.Printf("BUILDING ID: %d ", i.ID)
	now := time.Now()
	row := conn.QueryRow("UPDATE building SET name = ?, address = ?, longitude = ?, latitude = ?, updated_at = ? WHERE id = ?; ", i.Name, i.Address, i.Longitude, i.Latitude, now, i.ID)

	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem updating building ")
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
func DeleteBuilding(conn *sql.DB, buildingId string) error {
	row := conn.QueryRow("DELETE FROM building WHERE id = ?", buildingId)
	err := row.Scan()
	if err != sql.ErrNoRows {
		fmt.Println(" error deleting items %v", err)
		return err
	}
	return nil
}
