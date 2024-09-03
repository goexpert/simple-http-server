package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.Healthz)

	return mux
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
	log.Println("Healthz", r.RemoteAddr, r.RequestURI)
	resp := make(map[string]string)
	resp["message"] = "100%"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
