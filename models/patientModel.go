package models

import (
	"database/sql"
	"fmt"
	"golangCrud/config"
	"golangCrud/entities"
)

type PatientModel struct {
	conn *sql.DB
}

func NewPatientModel() *PatientModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)

	}

	return &PatientModel{
		conn: conn,
	}
}

func (p *PatientModel) Create(patient entities.Patient) bool {
	result, err := p.conn.Exec("insert into patient(name, nick, male_female, place_of_birth, date_of_birth, address, phone)value(?,?,?,?,?,?,?)",
		patient.Name, patient.Nick, patient.Male_female, patient.Place_of_birth, patient.Date_of_birth, patient.Address, patient.Phone)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
