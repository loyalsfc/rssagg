package main

import (
	"fmt"
	"net/http"

	"github.com/loyalsfc/rssagg/internal/auth"
	"github.com/loyalsfc/rssagg/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apicCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetAPIKey(r.Header)

		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		user, err := apicCfg.DB.GetUserByApiKey(r.Context(), apikey)

		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Auth Error: %v", err))
			return
		}

		handler(w, r, user)
	}
}
