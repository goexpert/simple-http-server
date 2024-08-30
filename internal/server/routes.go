package server

import (
	"encoding/json"
	"log"
	"net/http"
)

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.Healthz)
	mux.HandleFunc("GET /api/v1/blueprint", s.GetBlueprints)
	mux.HandleFunc("GET /api/v1/blueprint/{id}", s.GetBlueprintById)

	return mux
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "100%"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) GetBlueprints(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["blueprints"] = "100%"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (s *Server) GetBlueprintById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	resp := make(map[string]string)
	resp["blueprints"] = id
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)

}
