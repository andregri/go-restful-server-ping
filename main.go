package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/emicklei/go-restful"
)

func pingTime(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, fmt.Sprintf("%s\n", time.Now()))
}

func main() {
	// Create a web service
	webservice := new(restful.WebService)

	// Create a route and attach the handler to the service
	webservice.Route(webservice.GET("/ping").To(pingTime))

	// Add service to application
	restful.Add(webservice)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
