package handlers

import (
	"github.com/wunari/easypoll-backend/docs/models"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/poll"

	"github.com/go-openapi/runtime/middleware"
)

// in-memory "database", it will be changed to proper database later on
var pollTitles = []string{"Project Meeting", "Favorite Activities", "Favorite Food"}
var polls = []*models.Poll{
	{Title: &pollTitles[0], Slug: "project-meeting"},
	{Title: &pollTitles[1], Slug: "favotire-activities"},
	{Title: &pollTitles[2], Slug: "favorite-food"},
}

// GetPollsHandlerFunc returns an array of polls
func GetPollsHandlerFunc(params poll.GetPollsParams) middleware.Responder {
	return poll.NewGetPollsOK().WithPayload(polls)
}

// CreatePollHandlerFunc inserts a new poll in the database
func CreatePollHandlerFunc(params poll.CreatePollParams) middleware.Responder {
	newPoll := &models.Poll{Title: params.Body.Title, Slug: params.Body.Slug}
	polls = append(polls, newPoll)
	return poll.NewCreatePollOK().WithPayload(newPoll)
}
