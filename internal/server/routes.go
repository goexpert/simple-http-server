package server

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/goexpert/simple-http-server/internal/tracer"
	"go.opentelemetry.io/otel"
)

var oTracer = otel.Tracer("simple-http-server")

func (s *Server) RegisterRoutes() http.Handler {

	mux := http.NewServeMux()
	mux.HandleFunc("GET /", s.Healthz)

	return mux
}

func (s *Server) Healthz(w http.ResponseWriter, r *http.Request) {

	cleanup := tracer.InitTracer()
	defer cleanup()

	ctx, span := oTracer.Start(context.Background(), "Healthz")
	defer span.End()

	time.Sleep(time.Second * 1)

	sleep2(ctx)

	log.Println("Healthz", r.RemoteAddr, r.RequestURI)
	resp := make(map[string]string)
	resp["message"] = "100%"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func sleep2(ctx context.Context) {

	_, span := oTracer.Start(ctx, "sleep2")
	defer span.End()

	time.Sleep(time.Second * 2)

}
