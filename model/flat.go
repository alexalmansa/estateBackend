package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Flat struct {
	ID                int       `json:"id"`
	CreatedAt         time.Time `json:"_"`
	UpdatedAt         time.Time `json:"_"`
	BuildingId        int       `json:"building_id"`
	AskedPrice        int       `json:"asked_price"`
	Floor             int       `json:"floor"`
	DoorNumber        int       `json:"door_number"`
	Area              int       `json:"area"`
	BoilerDate        string    `json:"boiler_date"`
	BoilerDescription string    `json:"boiler_description"`
	PriceIndex        float32   `json:"price_index"`
}

func (i *Flat) Create(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("INSERT INTO flat (building_id, asked_price, floor, door_number, area ,boiler_date, boiler_description, price_index, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?,?,?)", i.BuildingId, i.AskedPrice, i.Floor, i.DoorNumber, i.Area, i.BoilerDate, i.BoilerDescription, i.PriceIndex, now, now)
	err := row.Scan(&i.ID, &i.BuildingId)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating Flat " + err.Error())
	}
	return nil

}

func (i *Flat) UpdateFlat(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("UPDATE flat SET asked_price = ?,floor = ?, door_number = ?, area = ?, updated_at = ?, building_id = ?, boiler_date = ?, boiler_description = ?, price_index = ? WHERE id = ?; ", i.AskedPrice, i.Floor, i.DoorNumber, i.Area, now, i.BuildingId, i.BoilerDate, i.BoilerDescription, i.PriceIndex, i.ID)

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
		rows, err = conn.Query("SELECT asked_price, floor, door_number, area, id, building_id, boiler_date, boiler_description, price_index FROM flat WHERE building_id = ? ORDER BY floor", buildingId)

	} else {
		rows, err = conn.Query("SELECT asked_price, floor, door_number, area, id, building_id, boiler_date, boiler_description, price_index FROM flat ORDER BY floor")
	}
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var flat []Flat
	for rows.Next() {
		i := Flat{}
		err = rows.Scan(&i.AskedPrice, &i.Floor, &i.DoorNumber, &i.Area, &i.ID, &i.BuildingId, &i.BoilerDate, &i.BoilerDescription, &i.PriceIndex)
		flat = append(flat, i)
	}
	return flat, nil
}

func GetFlat(conn *sql.DB, flatId string) (Flat, error) {
	var rows *sql.Rows
	var err error

	if flatId != "" {
		rows, err = conn.Query("SELECT asked_price, floor, door_number, area, id, building_id, boiler_date, boiler_description, price_index FROM flat WHERE id = ? ", flatId)

	} else {
		return Flat{}, err
	}
	if err != nil {
		fmt.Println(" error getting items %v", err)
		return Flat{}, err
	}
	var flat Flat
	for rows.Next() {
		i := Flat{}
		err = rows.Scan(&i.AskedPrice, &i.Floor, &i.DoorNumber, &i.Area, &i.ID, &i.BuildingId, &i.BoilerDate, &i.BoilerDescription, &i.PriceIndex)
		flat = i
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
