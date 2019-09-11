package main

import (
	"flag"
	"log"
	"os"
	"strconv"

	"github.com/wunari/easypoll-backend/docs/models"

	"github.com/wunari/easypoll-backend/docs/restapi/operations/auth"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/vote"
	"github.com/wunari/easypoll-backend/middleware"

	"github.com/wunari/easypoll-backend/docs/restapi"
	"github.com/wunari/easypoll-backend/docs/restapi/operations"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/poll"
	"github.com/wunari/easypoll-backend/handlers"

	"github.com/go-openapi/loads"
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

	// handlers
	api.PollGetPollsHandler = poll.GetPollsHandlerFunc(handlers.GetPollsHandlerFunc)
	api.PollCreatePollHandler = poll.CreatePollHandlerFunc(handlers.CreatePollHandlerFunc)
	api.PollGetPollByIDHandler = poll.GetPollByIDHandlerFunc(handlers.GetPollByIDHandlerFunc)
	api.PollUpdatePollByIDHandler = poll.UpdatePollByIDHandlerFunc(handlers.UpdatePollByIDHandlerFunc)
	api.PollDeletePollByIDHandler = poll.DeletePollByIDHandlerFunc(handlers.DeletePollByIDHandlerFunc)

	api.VoteAddVotePollHandler = vote.AddVotePollHandlerFunc(handlers.AddVotePollHandlerFunc)

	api.AuthGetAuthenticatedUserHandler = auth.GetAuthenticatedUserHandlerFunc(handlers.GetAuthenticatedUserHandlerFunc)

	// auth
	api.BearerAuth = func(token string) (*models.User, error) {
		return middleware.IsValidToken(token)
	}

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
