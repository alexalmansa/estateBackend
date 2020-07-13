package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Flat struct {
	ID        	int       `json:"id"`
	CreatedAt 	time.Time `json:"_"`
	UpdatedAt 	time.Time `json:"_"`
	BuildingId 	int       `json:"building_id"`
	RenterId  	int       `json:"renter_id"`
	Price     	int   	  `json:"price"`
}

func (i *Flat) Create(conn * pgx.Conn) error {

	now := time.Now()
	row := conn.QueryRow(context.Background(), "INSERT INTO flat (building_id, renter_id, price, created_at, updated_at) VALUES ($1, $2, $3, $4, $5) RETURNING id, building_id, renter_id", i.BuildingId, i.RenterId, i.Price, now, now)
	err := row.Scan(&i.ID, &i.BuildingId, &i.RenterId)
	if err != nil{
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating building")
	}
	return nil
}