package handler

import (
	"encoding/json"
	"fmt"
	"github.com/masseelch/render"
	"net/http"
	"poll/ent"
	usr "poll/ent/user"
	"strconv"
)

// Register creates a new user entry
func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var d ent.User
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		render.BadRequest(w, r, "invalid json string")
		return
	}
	user, err := h.Client.User.Create().
		SetEmail(d.Email).
		SetName(d.Name).
		SetPassword(d.Password).
		Save(r.Context())
	if err != nil {
		fmt.Printf("%v", err.Error())
		render.InternalServerError(w, r, "Failed to register the user")
		return
	}
	fmt.Println("User registered successfully")
	render.OK(w, r, user)
}


// Login verifies the login credentials of the user
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var d ent.User
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		render.BadRequest(w, r, "invalid json string")
		return
	}

	//Check if email exists
	user, err := h.Client.User.
		Query().
		Where(usr.Email(d.Email)).
		Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			render.NotFound(w, r, "Email Doesn't exists")
		case ent.IsNotSingular(err):
			render.BadRequest(w, r, "Invalid Email")
		default:
			render.InternalServerError(w, r, "Server Error")
		}
		return
	}

	// Verify the password
	if user.Password == d.Password {
		fmt.Println("User Verified. Log In Successful")
		render.OK(w, r, user)
		return
	}
	render.Unauthorized(w, r, "Invalid Email or Password.")
}


// GetUser returns the user details
func (h *Handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		render.BadRequest(w, r, "id must be an integer greater zero")
		return
	}

	// Query user details from userID
	user, err := h.Client.User.
		Query().
		Where(usr.ID(id)).
		Only(r.Context())
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			render.NotFound(w, r, "Email Doesn't exists")
		default:
			render.InternalServerError(w, r, "Server Error")
		}
		return
	}
	render.OK(w, r, user)
}
