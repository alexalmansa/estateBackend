package model

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"time"
)

type Flat struct {
	ID        	int       	`json:"id"`
	CreatedAt 	time.Time 	`json:"_"`
	UpdatedAt 	time.Time 	`json:"_"`
	BuildingId 	int			`json:"building_id"`
	AskedPrice  int			`json:"asked_price"`
	NumberDoor 	string 		`json:"number_door"`
	Area 		int 		`json:"area"`
}

func (i *Flat) Create(conn * pgx.Conn) error {

	now := time.Now()
	row := conn.QueryRow(context.Background(), "INSERT INTO flat (building_id, asked_price, number_door, area ,created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, building_id", i.BuildingId, i.AskedPrice, i.NumberDoor, i.Area, now, now)
	err := row.Scan(&i.ID, &i.BuildingId)
	if err != nil{
		fmt.Println(err)
		return fmt.Errorf("There was a problem creating building")
	}
	return nil
}