package main

import (
	"fmt"
	"github.com/olehhuss/rssagg/auth"
	"github.com/olehhuss/rssagg/internal/database"
	"net/http"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			respondWithError(w, 403, fmt.Sprintf("Cannot get api key: %v", err))
		}

		user, err := apiConfig.DB.GetUserByKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, 400, fmt.Sprintf("Cannot get user: %v", err))
			return
		}
		respondWithJSON(w, 200, databaseUserToUser(user))

		handler(w, r, user)
	}
}
