package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/xtasysensei/go-poll/internal/auth"
	"github.com/xtasysensei/go-poll/pkg/database"
	"github.com/xtasysensei/go-poll/pkg/models"
	"github.com/xtasysensei/go-poll/pkg/utils"
)

func HandleRegister(w http.ResponseWriter, r *http.Request) {
	var payload models.RegisterUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}
	//check if user exist
	_, err := GetUserByEmail(payload.Email, database.DB)
	if err == nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user with email %s already exists", payload.Email))
		return
	}
	//check if username already exists
	ok, err := IsUsernameTaken(database.DB, payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	if ok {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("username %s already exists", payload.Username))
		return
	}

	//check if passwords match
	if payload.Password != payload.ConfirmPassword {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("passwords dont match for %s", payload.Username))
		return
	}
	// if not, create user
	hashedPassword, err := auth.HashPassword(payload.Password)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	err = CreateUser(models.User{
		Username: payload.Username,
		Email:    payload.Email,
		Password: hashedPassword,
	}, database.DB)

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusCreated, map[string]string{"message": "user successfully created"})
}
