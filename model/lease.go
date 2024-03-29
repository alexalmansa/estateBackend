package model

import (
	"database/sql"
	"estateBackend/utils"
	"fmt"
	"time"
)

type Lease struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"_"`
	UpdatedAt time.Time `json:"_"`
	FlatId    int       `json:"flat_id"`
	RenterId  int       `json:"renter_id"`
	Price     float64   `json:"price"`
	StartDate string    `json:"start_date"`
	EndDate   string    `json:"end_date"`
	Deposit   float64   `json:"deposit"`
}

func (i *Lease) Create(conn *sql.DB) error {

	now := time.Now()
	startDate, errStart := utils.ConvertTime(i.StartDate)
	endDate, errEnd := utils.ConvertTime(i.EndDate)
	if startDate.After(endDate) {
		return fmt.Errorf("Error end date is before beggining date")
	}
	if errStart != nil || errEnd != nil {
		fmt.Println("Error with date format")
	} else {
		row := conn.QueryRow("INSERT INTO lease (flat_id, renter_id, price, start_date ,end_date, deposit, created_at, updated_at) VALUES (?,?,?,?,?,?,?,?)", i.FlatId, i.RenterId, i.Price, startDate, endDate, i.Deposit, now, now)
		err2 := row.Scan(&i.ID)
		if err2 != sql.ErrNoRows {
			fmt.Println(err2)
			return fmt.Errorf("There was a problem creating lease")
		}
	}
	return nil

}

func (i *Lease) GetAllLeases(conn *sql.DB, flatId string, renterId string, pastLeases string) ([]Lease, error) {
	var rows *sql.Rows
	var err error

	if renterId != "" && flatId != "" {
		if pastLeases == "false" {
			rows, err = conn.Query("SELECT id, end_date, start_date, deposit, price, renter_id, flat_id FROM lease WHERE renter_id = ? AND flat_id = ? AND end_date >= end_date >= current_date() ORDER BY end_date DESC ", renterId, flatId)
		} else {
			rows, err = conn.Query("SELECT id, end_date, start_date, deposit, price, renter_id, flat_id FROM lease WHERE renter_id = ? AND flat_id = ? ORDER BY end_date DESC", renterId, flatId)
		}
	} else if flatId != "" {
		if pastLeases == "false" {
			rows, err = conn.Query("SELECT id, end_date, start_date, deposit, price, renter_id, flat_id FROM lease WHERE flat_id = ? AND end_date >= end_date >= current_date() ORDER BY end_date DESC", flatId)
		} else {
			rows, err = conn.Query("SELECT id, end_date, start_date, deposit, price, renter_id, flat_id FROM lease WHERE flat_id = ? ORDER BY end_date DESC", flatId)
		}
	} else if renterId != "" {

		rows, err = conn.Query("SELECT id, end_date, start_date, deposit, price, renter_id, flat_id FROM lease WHERE renter_id = ? ORDER BY end_date DESC", renterId)

	} else {
		if pastLeases == "false" {

			rows, err = conn.Query("SELECT id, end_date, start_date, deposit, price, renter_id, flat_id FROM lease WHERE end_date >= current_date() ORDER BY end_date DESC")
		} else {
			rows, err = conn.Query("SELECT id, end_date, start_date, deposit, price, renter_id, flat_id FROM lease ORDER BY end_date DESC")

		}
	}

	if err != nil {
		fmt.Println(" error getting items %v", err)
		return nil, err
	}
	var lease []Lease
	for rows.Next() {
		i := Lease{}
		err = rows.Scan(&i.ID, &i.EndDate, &i.StartDate, &i.Deposit, &i.Price, &i.RenterId, &i.FlatId)
		lease = append(lease, i)
	}
	return lease, nil
}

func (i *Lease) DeleteLease(conn *sql.DB, leaseId string) error {
	row := conn.QueryRow("DELETE FROM lease WHERE id = ?", leaseId)
	err := row.Scan()
	if err != sql.ErrNoRows {
		fmt.Println(" error deleting items %v", err)
		return err
	}
	return nil
}

func (i *Lease) UpdateLease(conn *sql.DB) error {

	now := time.Now()
	row := conn.QueryRow("UPDATE lease SET end_date = ?, start_date = ?, deposit = ?, price = ?, renter_id = ?, flat_id = ?, updated_at = ? WHERE id = ?; ", i.EndDate, i.StartDate, i.Deposit, i.Price, i.RenterId, i.FlatId, now, i.ID)

	err := row.Scan(&i.ID)
	if err != sql.ErrNoRows {
		fmt.Println(err)
		return fmt.Errorf("There was a problem updating lease ")
	}
	return nil
}
