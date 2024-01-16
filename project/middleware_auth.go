package main

import (
	"fmt"
	"net/http"

	"github.com/cdrojasm/golang_freeCodeCamp/project/internal/auth"
	"github.com/cdrojasm/golang_freeCodeCamp/project/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			responseWithError(w, 403, fmt.Sprintf("Couldn't log with provided credentials: %s", err))
			return
		}
		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)

		if err != nil {
			responseWithError(w, 400, fmt.Sprintf("Couldn't get user: %s", err))
			return
		}
		handler(w, r, user)
	}
}
