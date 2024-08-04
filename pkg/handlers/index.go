package handlers

import (
	"net/http"

	"github.com/xtasysensei/go-poll/pkg/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "welcome to the go auth app"})
}
