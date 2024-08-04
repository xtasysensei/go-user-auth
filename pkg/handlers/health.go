package handlers

import (
	"net/http"

	"github.com/xtasysensei/go-poll/pkg/utils"
)

func Health(w http.ResponseWriter, r *http.Request) {
	utils.WriteJSON(w, http.StatusOK, map[string]string{"message": "Server is up and running"})
}
