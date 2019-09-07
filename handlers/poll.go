package handlers

import (
	"time"

	"github.com/wunari/easypoll-backend/docs/models"
	"github.com/wunari/easypoll-backend/docs/restapi/operations/poll"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
)

// in-memory "database", it will be changed to proper database later on
var polls = []*models.Poll{
	{ID: 1, Question: "What's your favorite color?", Answers: []*models.Answer{&models.Answer{Title: "Red"}, &models.Answer{Title: "Blue"}}, MultipleAnswers: false, CreatedAt: strfmt.DateTime(time.Now())},
	{ID: 2, Question: "What's your favorite fruit?", Answers: []*models.Answer{&models.Answer{Title: "Apple"}, &models.Answer{Title: "Orange"}}, MultipleAnswers: true, CreatedAt: strfmt.DateTime(time.Now())},
	{ID: 3, Question: "Do you like my car?", Answers: []*models.Answer{&models.Answer{Title: "Yes"}, &models.Answer{Title: "No"}}, MultipleAnswers: false, CreatedAt: strfmt.DateTime(time.Now())},
}
var pollCount = float64(len(polls))

// GetPollsHandlerFunc returns an array of polls
func GetPollsHandlerFunc(params poll.GetPollsParams) middleware.Responder {
	return poll.NewGetPollsOK().WithPayload(polls)
}

// CreatePollHandlerFunc inserts a new poll in the database
func CreatePollHandlerFunc(params poll.CreatePollParams) middleware.Responder {
	pollCount++
	newPoll := &models.Poll{
		ID:              pollCount,
		Question:        params.Body.Question,
		Answers:         params.Body.Answers,
		MultipleAnswers: params.Body.MultipleAnswers,
		CreatedAt:       strfmt.DateTime(time.Now()),
	}
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
			polls[i] = &models.Poll{
				ID:              p.ID,
				Question:        params.Body.Question,
				Answers:         params.Body.Answers,
				MultipleAnswers: params.Body.MultipleAnswers,
				CreatedAt:       p.CreatedAt,
			}
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
