package app

import (
	"encoding/json"
	"net/http"
)

type Server struct {
	name string
}

func New(name string) *Server {
	return &Server{name: name}
}

func (s *Server) Handler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", s.health)
	mux.HandleFunc("GET /ready", s.ready)
	return mux
}

func (s *Server) health(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, http.StatusOK, map[string]any{
		"status":  "healthy",
		"ready":   true,
		"service": s.name,
	})
}

func (s *Server) ready(writer http.ResponseWriter, _ *http.Request) {
	writeJSON(writer, http.StatusOK, map[string]any{
		"ready": true,
	})
}

func writeJSON(writer http.ResponseWriter, status int, body map[string]any) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	if err := json.NewEncoder(writer).Encode(body); err != nil {
		http.Error(writer, "encode response", http.StatusInternalServerError)
	}
}
