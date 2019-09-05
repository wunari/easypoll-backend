package handlers

import (
	"github.com/wunari/easypoll-backend/docs/models"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/poll"

	"github.com/go-openapi/runtime/middleware"
)

// in-memory "database", it will be changed to proper database later on
var pollTitles = []string{"Project Meeting", "Favorite Activities", "Favorite Food"}
var polls = []*models.Poll{
	{ID: 1, Title: &pollTitles[0], Slug: "project-meeting"},
	{ID: 2, Title: &pollTitles[1], Slug: "favotire-activities"},
	{ID: 3, Title: &pollTitles[2], Slug: "favorite-food"},
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

// GetPollByIDHandlerFunc gets a single poll by id
func GetPollByIDHandlerFunc(params poll.GetPollByIDParams) middleware.Responder {
	for _, p := range polls {
		if p.ID == params.ID {
			return poll.NewGetPollByIDOK().WithPayload(p)
		}
	}
	return poll.NewGetPollByIDNotFound()
}

// UpdatePollByIDHandlerFunc updates a poll by id
func UpdatePollByIDHandlerFunc(params poll.UpdatePollByIDParams) middleware.Responder {
	for i, p := range polls {
		if p.ID == params.ID {
			polls[i] = &models.Poll{ID: params.ID, Title: params.Body.Title, Slug: params.Body.Slug}
			return poll.NewUpdatePollByIDOK().WithPayload(polls[i])
		}
	}
	return poll.NewUpdatePollByIDNotFound()
}

// DeletePollByIDHandlerFunc deletes a poll by id
func DeletePollByIDHandlerFunc(params poll.DeletePollByIDParams) middleware.Responder {
	for i, p := range polls {
		if p.ID == params.ID {
			polls = polls[:i+copy(polls[i:], polls[i+1:])]
			return poll.NewDeletePollByIDNoContent()
		}
	}
	return poll.NewDeletePollByIDNotFound()
}
