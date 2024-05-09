package api

import (
	"database/sql"
	"log"
	"net/http"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		// db:   db,
	}
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("GET /test/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		w.Write([]byte("Id: " + id))
	})

	v1 := http.NewServeMux()
	v1.Handle("/api/v1/", http.StripPrefix("/api/v1", router))

	middlewareChain := MiddlewareChain(
		RequestLoggerMiddleware,
		// RequireAuthMiddleware
	)
	server := http.Server{
		Addr:    s.addr,
		Handler: middlewareChain(router),
	}

	return server.ListenAndServe()
}

func RequestLoggerMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("method %s, path: %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}

func RequireAuthMiddleware(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// check authentication
		token := r.Header.Get("Authorization")
		if token != "Bearer Token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	}
}

type Middleware func(http.Handler) http.HandlerFunc

func MiddlewareChain(middlwares ...Middleware) Middleware {
	return func(next http.Handler) http.HandlerFunc {
		for i := len(middlwares) - 1; i >= 0; i-- {
			next = middlwares[i](next)
		}

		return next.ServeHTTP
	}
}
