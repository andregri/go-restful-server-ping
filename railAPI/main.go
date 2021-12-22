package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/andregri/go-restful-server-ping/dbutils"
	"github.com/emicklei/go-restful"
	_ "github.com/mattn/go-sqlite3"
)

// DB driver is a global variable visible to the whole program
var DB *sql.DB

// TrainResoure is the model for holding train information
type TrainResource struct {
	ID              int
	DriverName      string
	OperatingStatus bool
}

// StationResource holds information about locations
type StationResource struct {
	ID          int
	Name        string
	OpeningTime time.Time
	ClosingTime time.Time
}

// RouteResource links both trains and stations
type RouteResource struct {
	ID          int
	TrainID     int
	StationID   int
	ArrivalTime time.Time
}

// Register add paths and routes to container
func (t *TrainResource) Register(container *restful.Container) {
	ws := new(restful.WebService)
	ws.
		Path("/v1/trains").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)

	ws.Route(ws.GET("/{train-id}").To(t.getTrain))
	ws.Route(ws.POST("").To(t.createTrain))
	ws.Route(ws.DELETE("/{train-id}").To(t.deleteTrain))

	container.Add(ws)
}

// GET http://localhost:8000/v1/trains/1
func (t TrainResource) getTrain(req *restful.Request, res *restful.Response) {
	// Get train-id parameter from url
	id := req.PathParameter("train-id")

	// find train by id in the db
	err := DB.QueryRow(`
		SELECT id, driver_name, operating_status FROM train
		WHERE id=?
	`, id).Scan(&t.ID, &t.DriverName, &t.OperatingStatus)
	if err != nil {
		log.Println("GET trains:", err)
		res.AddHeader("Content-Type", "text/plain")
		res.WriteErrorString(http.StatusNotFound, "Train could not be found")
	} else {
		res.WriteEntity(t)
	}
}

// POST http://localhost:8000/v1/trains
func (t TrainResource) createTrain(req *restful.Request, res *restful.Response) {
	// Decode input from json
	log.Println(req.Request.Body)
	decoder := json.NewDecoder(req.Request.Body)
	var tempTrain TrainResource
	err := decoder.Decode(&tempTrain)
	if err != nil {
		log.Println("json decode error:", err)
	}

	// Create insert statement
	statement, err := DB.Prepare(`
		INSERT INTO train (driver_name, operating_status)
		VALUES (?, ?)
	`)
	if err != nil {
		log.Println("insert statement error:", err)
	}

	// Create row in db
	result, err := statement.Exec(tempTrain.DriverName, tempTrain.OperatingStatus)
	if err != nil {
		res.AddHeader("Content-Type", "text/plain")
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		newID, _ := result.LastInsertId()
		tempTrain.ID = int(newID)
		res.WriteHeaderAndEntity(http.StatusCreated, tempTrain)
	}
}

// DELETE http://localhost:8000/v1/trains/1
func (t TrainResource) deleteTrain(req *restful.Request, res *restful.Response) {
	// read id from path parameters
	id := req.PathParameter("train-id")

	// Create statement for deleting a resource from db
	statement, err := DB.Prepare("DELETE FROM train WHERE id=?")
	if err != nil {
		log.Println("delete statement error:", err)
	}

	// Delete resource
	_, err = statement.Exec(id)

	// Write result to response
	if err != nil {
		res.AddHeader("Content-Type", "text/plain")
		res.WriteErrorString(http.StatusInternalServerError, err.Error())
	} else {
		res.WriteHeader(http.StatusOK)
	}
}

// To ping API server and check it is alive
func pingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, fmt.Sprintf("%s\n", time.Now()))
}

func main() {
	// Connect to database
	var err error
	DB, err = sql.Open("sqlite3", "./railapi.db")
	if err != nil {
		log.Println("Driver creation failed")
	}

	// Create tables
	dbutils.Initialize(DB)

	// Create a restful container using a "curly router",
	// using {param} syntax for parameters in routes
	wsContainer := restful.NewContainer()
	wsContainer.Router(restful.CurlyRouter{})

	// Handle routing for /v1/trains
	t := TrainResource{}
	t.Register(wsContainer)

	// Create a web service for /ping
	wsPing := new(restful.WebService)
	wsPing.Route(wsPing.GET("/ping").To(pingTime))
	wsContainer.Add(wsPing)

	// Start the server
	server := &http.Server{
		Addr:    ":8000",
		Handler: wsContainer,
	}
	log.Fatal(server.ListenAndServe())
}
