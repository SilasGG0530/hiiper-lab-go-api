package main

import (
	_ "bytes"
	_ "encoding/json"
	"hiiper-lab-go-api"
	"log"
	_ "net/http"
	_ "net/http/httptest"
	"os"
	_ "strconv"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM offline_measurements_new")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS offline_measurements_new
(
	id text NOT NULL,
	"name" text NULL,
	price varchar NULL,
	confirmed_by text NOT NULL
)`
