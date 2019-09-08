package handlers

import (
	"github.com/wunari/easypoll-backend/docs/restapi/operations/vote"

	"github.com/go-openapi/runtime/middleware"
)

// AddVotePollHandlerFunc adds a vote to an answer
func AddVotePollHandlerFunc(params vote.AddVotePollParams) middleware.Responder {
	foundPoll := false
	votedAnswers := make([]int, 10)
	for i, poll := range Polls {
		if poll.ID == params.ID {
			if !poll.MultipleAnswers && len(params.Body) > 1 {
				response := &vote.AddVotePollBadRequestBody{Code: 400, Message: "This poll doesn't accept multiple answers"}
				return vote.NewAddVotePollBadRequest().WithPayload(response)
			} else if len(params.Body) < 1 {
				response := &vote.AddVotePollBadRequestBody{Code: 400, Message: "You need to select at least one answer"}
				return vote.NewAddVotePollBadRequest().WithPayload(response)
			}

			for j := range poll.Answers {
				for _, k := range params.Body {
					if int(k) == j && !contains(votedAnswers, j) {
						Polls[i].Votes++
						Polls[i].Answers[j].Votes++

						foundPoll = true
						votedAnswers = append(votedAnswers, j)
					}
				}
			}
		}
	}

	if foundPoll == false {
		return vote.NewAddVotePollNotFound()
	}
	return vote.NewAddVotePollNoContent()
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
