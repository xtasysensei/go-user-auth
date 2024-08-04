package user

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/xtasysensei/go-poll/internal/auth"
	"github.com/xtasysensei/go-poll/internal/config"
	"github.com/xtasysensei/go-poll/pkg/database"
	"github.com/xtasysensei/go-poll/pkg/models"
	"github.com/xtasysensei/go-poll/pkg/utils"
)

func HandleLogin(w http.ResponseWriter, r *http.Request) {
	//get json payload

	var payload models.LoginUserPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
		return
	}

	//validate payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	//get user from database by username
	u, err := GetUserByUsername(database.DB, payload.Username)
	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("user does not exist"))
		return
	}

	//check password
	if !auth.ComparePasswords(u.Password, payload.Password) {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("passwords does not match"))
		return
	}
	secret := []byte(config.Envs.JWTSecret)
	token, err := auth.CreateJWT(secret, u.UserID)
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Login successful", "token": token})

}
