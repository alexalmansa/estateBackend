package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Flat struct {
	ID         int       `json:"id"`
	CreatedAt  time.Time `json:"_"`
	UpdatedAt  time.Time `json:"_"`
	BuildingId int       `json:"building_id"`
	AskedPrice int       `json:"asked_price"`
	NumberDoor string    `json:"number_door"`
	Area       int       `json:"area"`
}

func (i *Flat) Create(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("INSERT INTO flat (building_id, asked_price, number_door, area ,created_at, updated_at) VALUES (?,?,?,?,?,?)", i.BuildingId, i.AskedPrice, i.NumberDoor, i.Area, now, now)
	err := row.Scan(&i.ID, &i.BuildingId)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating Flat " + err.Error())
	}
	return nil

}

func (i *Flat) UpdateFlat(conn *sql.DB) error {

	fmt.Printf("BUILDING ID: %d ", i.ID)
	now := time.Now()
	row := conn.QueryRow("UPDATE flat SET asked_price = ?, number_door = ?, area = ?, updated_at = ?, building_id = ? WHERE id = ?; ", i.AskedPrice, i.NumberDoor, i.Area, now, i.BuildingId, i.ID)

	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem updating flat ")
	}
	return nil
}

func GetBuildingItems(conn *sql.DB, buildingId string) ([]Flat, error) {
	var rows *sql.Rows
	var err error

	if buildingId != "" {
		rows, err = conn.Query("SELECT asked_price, number_door, area, id, building_id FROM flat WHERE building_id = ?", buildingId)

	} else {
		rows, err = conn.Query("SELECT asked_price, number_door, area, id, building_id FROM flat")
	}
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var flat []Flat
	for rows.Next() {
		i := Flat{}
		err = rows.Scan(&i.AskedPrice, &i.NumberDoor, &i.Area, &i.ID, &i.BuildingId)
		flat = append(flat, i)
	}
	return flat, nil
}

func DeleteFlat(conn *sql.DB, flatId string) error {
	row := conn.QueryRow("DELETE FROM flat WHERE id = ?", flatId)
	err := row.Scan()
	if err != sql.ErrNoRows {
		fmt.Println(" error deleting items %v", err)
		return err
	}
	return nil
}
