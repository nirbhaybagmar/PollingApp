package handler

import (
	"encoding/json"
	"github.com/masseelch/render"
	"net/http"
	"poll/ent"
	"poll/ent/question"
	"strconv"
)

type Poll struct {
	Question ent.Question `json:"question,omitempty"`
	Answers  ent.Answers  `json:"answers,omitempty"`
}

// GetAllPolls calls request to retrieve questions along with its answer
func (h *Handler) GetAllPolls(w http.ResponseWriter, r *http.Request) {
	polls, err := h.Client.Question.
		Query().
		WithAnswer().
		WithUser().
		All(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Error Retrieving polls")
		return
	}
	render.OK(w, r, polls)
}

// GetAllPollsByUser calls request to retrieve questions along with its answer for specific user
func (h *Handler) GetAllPollsByUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}

	polls, err := h.Client.Question.
		Query().
		Where(question.UserID(id)).
		WithAnswer().
		All(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Error Retrieving polls")
		return
	}
	render.OK(w, r, polls)
}

// CreateNewPoll creates a new poll along with entries for option
func (h *Handler) CreateNewPoll(w http.ResponseWriter, r *http.Request) {
	var d Poll
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		render.BadRequest(w, r, "invalid json string")
		return
	}
	// Creating a new question
	ques, err := h.Client.Question.
		Create().
		SetBody(d.Question.Body).
		SetUserID(d.Question.UserID).
		SetPollExpiry(d.Question.PollExpiry).
		Save(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Failed to create poll")
		return
	}

	// Creating multiple answers along with its Question
	bulk := make([]*ent.AnswerCreate, len(d.Answers))
	for i, name := range d.Answers {
		bulk[i] = h.Client.Answer.Create().
			SetBody(name.Body).
			AddQuestionID(ques)
	}

	_, err = h.Client.Answer.CreateBulk(bulk...).Save(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Failed to create poll")
		return
	}

	// Render OK status
	render.OK(w, r, "Poll created successfully.")
}

// GetAnswers returns a list of all the answers with its question id
func (h *Handler) GetAnswers(w http.ResponseWriter, r *http.Request) {
	ans, err := h.Client.Answer.Query().
		WithQuestionID().
		All(r.Context())
	if err != nil {
		render.InternalServerError(w, r, "Error Retrieving polls")
		return
	}
	render.OK(w, r, ans)
}
