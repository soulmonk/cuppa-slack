package protocol

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"os"
)

type response struct {
	Success bool `json:"success"`
}

func Rest() {
	port := os.Getenv("REST_PORT")
	if port == "" {
		port = ":8080"
	}
	http.HandleFunc("/ping", ping)

	log.Debug().Str("port", port).Msg("Server starting...")
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal().Err(err).Msg("can not start server")
	}
}

// Handler
func ping(w http.ResponseWriter, r *http.Request) {
	data := response{true}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Error().Err(err).Msg("can not send response")
	}
}
