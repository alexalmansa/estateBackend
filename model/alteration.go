package model

import (
	"database/sql"
	"estateBackend/utils"
	"fmt"
	"time"
)

type Alteration struct {
	ID          int       `json:"id"`
	CreatedAt   time.Time `json:"_"`
	UpdatedAt   time.Time `json:"_"`
	FlatId      int       `json:"flat_id"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
	Date        string    `json:"alter_date"`
}

func (i *Alteration) Create(conn *sql.DB) error {

	now := time.Now()
	date, errDate := utils.ConvertTime(i.Date)

	if errDate != nil {
		fmt.Println("Error with date format")
	} else {
		row := conn.QueryRow("INSERT INTO alterations (flat_id, price, alter_date, created_at, updated_at) VALUES (?,?,?,?,?)", i.FlatId, i.Price, date, now, now)
		err2 := row.Scan(&i.ID)
		if err2 != sql.ErrNoRows {
			fmt.Println(err2)
			return fmt.Errorf("There was a problem creating lease")
		}
	}
	return nil

}

func GetAlterations(conn *sql.DB, flatId string) ([]Alteration, error) {
	var rows *sql.Rows
	var err error

	if flatId != "" {

		rows, err = conn.Query("SELECT id,flat_id, price, alter_date FROM alterations WHERE flat_id = ? ", flatId)

	} else {
		rows, err = conn.Query("SELECT id,flat_id, price, alter_date FROM alterations")

	}

	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var alteration []Alteration
	for rows.Next() {
		i := Alteration{}
		err = rows.Scan(&i.ID, &i.FlatId, &i.Price, &i.Date)
		alteration = append(alteration, i)
	}
	return alteration, nil
}

func DeleteAlteration(conn *sql.DB, alterationId string) error {
	row := conn.QueryRow("DELETE FROM alterations WHERE id = ?", alterationId)
	err := row.Scan()
	if err != sql.ErrNoRows {
		fmt.Println(" error deleting items %v", err)
		return err
	}
	return nil
}

func (i *Alteration) UpdateAlteration(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("UPDATE alterations SET flat_id = ?, price = ?, alter_date  = ?, updated_at = ? WHERE id = ?; ", i.FlatId, i.Price, i.Date, now, i.ID)

	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem updating alteration ")
	}
	return nil
}
