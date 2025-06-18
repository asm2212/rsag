package main

import (
	"github.com/asm2212/rsag/internal/auth"
	"github.com/asm2212/rsag/internal/database"
	"github.com/asm2212/rsag/internal/handlers"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, *database.User)

func (apiCfg *apiConfig) authMiddleware(next authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithError(w, 403, "error getting API key")
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, "could not get user")
			return
		}

		next(w, r, &user)
	}
}
