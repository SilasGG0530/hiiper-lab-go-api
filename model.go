package main

import (
	"database/sql"
	_ "database/sql"
	"errors"
	_ "errors"
)

type measurements struct {
	id           string `json:"id"`
	name         string `json:"name"`
	price        string `json:"price"`
	confirmed_by string `json:"confirmed_by"`
}

func (p *measurements) getmeasurements(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *measurements) updatemeasurements(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *measurements) deletemeasurements(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (p *measurements) createmeasurements(db *sql.DB) error {
	return errors.New("Not implemented")
}

func getmeasurementss(db *sql.DB, start, count int) ([]measurements, error) {
	return nil, errors.New("Not implemented")
}
