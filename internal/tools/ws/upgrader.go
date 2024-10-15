package ws

import (
	"net/http"

	"github.com/gorilla/websocket"
)

func GetUpgrader(allowedOrigins []string) *websocket.Upgrader {
	return &websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			if allowedOrigins == nil {
				return true
			}

			origin := r.Header.Get("Origin")
			for _, allowedOrigin := range allowedOrigins {
				if origin == allowedOrigin {
					return true
				}
			}
			return false
		},
		Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
			// do nothing
		},
	}
}
