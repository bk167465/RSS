package main

import (
	"fmt"
	"net/http"

	"github.com/bk167465/rss/internal/auth"
	"github.com/bk167465/rss/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        apiKey, err := auth.GetApiKey(r.Header)
	    if err != nil {
	    	responseWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
	    }
    
	    user, err := apiCfg.DB.GetUserAPIKey(r.Context(), apiKey)
	    if err != nil {
	    	responseWithError(w, 400, fmt.Sprintf("Coldn't get user: %v", err))
	    	return
	    }
    
	    handler(w, r, user)
	}
}