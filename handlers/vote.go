package handlers

import (
	"github.com/wunari/easypoll-backend/docs/restapi/operations/vote"

	"github.com/go-openapi/runtime/middleware"
)

// AddVotePollHandlerFunc adds a vote to an answer
func AddVotePollHandlerFunc(params vote.AddVotePollParams) middleware.Responder {
	for i, poll := range Polls {
		if poll.ID == params.ID {
			for j := range poll.Answers {
				if j == int(params.AnswerID) {
					Polls[i].Votes++
					Polls[i].Answers[j].Votes++
					return vote.NewAddVotePollNoContent()
				}
			}
		}
	}
	return vote.NewAddVotePollNotFound()
}
