package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/wunari/easypoll-backend/docs/models"
	"github.com/wunari/easypoll-backend/docs/restapi"
	"github.com/wunari/easypoll-backend/docs/restapi/operations"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/poll"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
)

func main() {
	// load port
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	if port == 0 {
		port = 3000
	}
	var portFlag = flag.Int("port", port, "Port to run this service on")

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewEasypollAPI(swaggerSpec)
	server := restapi.NewServer(api)
	server.EnabledListeners = []string{"http"}
	defer func() {
		_ = server.Shutdown()
	}()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	// GetGreetingHandler greets the given name,
	// in case the name is not given, it will default to World
	api.PollGetPollsHandler = poll.GetPollsHandlerFunc(
		func(params poll.GetPollsParams) middleware.Responder {
			polls := []*models.Poll{}
			return poll.NewGetPollsOK().WithPayload(polls)
		})

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
