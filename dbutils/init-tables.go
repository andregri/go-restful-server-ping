package dbutils

import (
	"database/sql"
	"log"
)

func createTable(dbDriver *sql.DB, table string) {
	statement, err := dbDriver.Prepare(table)
	if err != nil {
		log.Println(err)
	} else {
		_, err = statement.Exec()
		if err != nil {
			log.Println("Table already exists")
		}
	}
}

func Initialize(dbDriver *sql.DB) {
	createTable(dbDriver, trainTable)
	createTable(dbDriver, stationTable)
	createTable(dbDriver, routeTable)
	log.Println("All tables have been created")
}
