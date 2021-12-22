package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/andregri/go-restful-server-ping/dbutils"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

func pingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, fmt.Sprintf("%s\n", time.Now()))
}

func main() {
	// Connect to database
	db, err := sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed")
	}

	// Create tables
	dbutils.Initialize(db)

	// Create a web service
	webservice := new(restful.WebService)

	// Create a route and attach the handler to the service
	webservice.Route(webservice.GET("/ping").To(pingTime))

	// Add service to application
	restful.Add(webservice)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
