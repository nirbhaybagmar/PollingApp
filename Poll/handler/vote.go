package handler

import (
	"encoding/json"
	"fmt"
	"github.com/masseelch/render"
	"net/http"
	"poll/ent"
	ans "poll/ent/answer"
	"poll/ent/vote"
	"strconv"
)

// CastVote handles the vote of the poll
func (h *Handler) CastVote(w http.ResponseWriter, r *http.Request) {
	var d ent.Vote
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		render.BadRequest(w, r, "invalid json string")
		return
	}

	vote, err := h.Client.Vote.Create().
		SetAnswerID(d.AnswerID).
		SetQuestionID(d.QuestionID).
		SetUserID(d.UserID).Save(r.Context())

	if err != nil {
		render.InternalServerError(w, r, "Failed to register the vote")
		return
	}

	// Update the total vote count for the answer
	answer, err := h.Client.Answer.Query().
		Where(ans.ID(vote.AnswerID)).
		First(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Failed to update the answer count")
		return
	}

	_, _ = h.Client.Answer.
		UpdateOneID(vote.AnswerID).
		SetNumOfVotes(answer.NumOfVotes + 1).
		Save(r.Context())

	fmt.Printf("Registered Vote for question id: %v", vote.QuestionID)
	render.OK(w, r, "Vote Registered successfully.")
}

// GetVotes returns the list of all the votes
func (h *Handler) GetVotes(w http.ResponseWriter, r *http.Request) {
	ans, err := h.Client.Vote.Query().
		WithQuestion().
		WithAnswer().
		All(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Error Retrieving votes")
		return
	}
	render.OK(w, r, ans)
}

// GetVotesByUser returns the list of all the votes made by user.
func (h *Handler) GetVotesByUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}

	res, err := h.Client.Vote.Query().
		Select("question_id", "answer_id", "user_id").
		Where(vote.UserID(id)).
		All(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Error Retrieving votes")
		return
	}
	render.OK(w, r, res)
}
