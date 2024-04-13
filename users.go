package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

type UserService struct {
	store Store
}

func NewUserService(s Store) *UserService {
	return &UserService{store: s}
}

// reg the page register by handler handleUserRegister
func (s *UserService) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/users/register", s.handleUserRegister).Methods("POST")
}

func (s *UserService) handleUserRegister(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "error reading req body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	var payload *User
	err = json.Unmarshal(body, &payload)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: "Invalid req payload"})
		return
	}

	// check if any of the fields are empty
	if err := validateUserPayload(payload); err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}

	// store hashed pw
	hashedPW, err := HashPassword(payload.Password)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}
	payload.Password = hashedPW

	u, err := s.store.CreateUser(payload)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating user"})
		return
	}

	token, err := createAndSetAuthCookie(u.ID, w)
	if err != nil {
		WriteJSON(w, http.StatusInternalServerError, ErrorResponse{Error: "Error creating session"})
		return
	}

	WriteJSON(w, http.StatusCreated, token)

}

func validateUserPayload(user *User) error {
	if user.Email == "" {
		return errors.New("email is required")
	}

	if user.FirstName == "" {
		return errors.New("first name is required")
	}

	if user.LastName == "" {
		return errors.New("last name is required")
	}

	if user.Password == "" {
		return errors.New("password is required")
	}

	return nil
}

func createAndSetAuthCookie(id int64, w http.ResponseWriter) (string, error) {
	secret := []byte(Envs.JWTSecret)
	token, err := CreateJWT(secret, id)
	if err != nil {
		return "", err
	}
	http.SetCookie(w, &http.Cookie{
		Name:  "Authorization",
		Value: token,
	})
	return token, nil
}
