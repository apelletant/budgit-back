package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/apelletant/budgit/pkg/domain"

	"github.com/rs/cors"

	Log "github.com/apelletant/logger"

	"golang.org/x/sync/errgroup"
)

var (
	ErrMissingServer = errors.New("server should be initialized")
	ErrMissingLogger = errors.New("logger should be initialized")
)

type Config struct {
	Port int
}

func (cfg *Config) validate() error {
	return nil
}

type Dependencies struct {
	App domain.App
	Log *Log.Logger
}

func (d *Dependencies) validate() error {
	if d.Log == nil {
		return fmt.Errorf("dependencies.Log: %w", ErrMissingLogger)
	}

	if d.App == nil {
		return fmt.Errorf("dependencies.App: %w", ErrMissingServer)
	}

	return nil
}

type Server struct {
	cfg  *Config
	deps *Dependencies
	s    *http.Server
}

func New(deps *Dependencies, cfg *Config) (*Server, error) {
	if err := cfg.validate(); err != nil {
		return nil, fmt.Errorf("cfg.validate: %w", err)
	}

	if err := deps.validate(); err != nil {
		return nil, fmt.Errorf("deps.validate: %w", err)
	}

	srv := &Server{
		cfg:  cfg,
		deps: deps,
	}

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedMethods:   []string{"GET", "UPDATE", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token", "Cookie"},
		AllowCredentials: true,
		Debug:            true,
	})

	mux := http.NewServeMux()
	mux.Handle("POST /expense", http.HandlerFunc(srv.AddExpense))
	mux.Handle("GET /expenses", http.HandlerFunc(srv.GetAllExpenses))

	srv.s = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: c.Handler(mux),
	}

	return srv, nil
}

func (srv *Server) Run(ctx context.Context) error {

	errG, errCtx := errgroup.WithContext(ctx)

	errG.Go(func() error {
		srv.deps.Log.Info("starting server on port:", srv.cfg.Port)

		err := srv.s.ListenAndServe()
		if err != nil {
			return fmt.Errorf("h.srv.ListenAndServe: %w", err)
		}

		return nil
	})

	<-errCtx.Done()

	if err := srv.s.Close(); err != nil {
		srv.deps.Log.Error("h.srv.Close", err)
	}

	if err := errG.Wait(); err != nil {
		return fmt.Errorf("errG.Wait: %w", err)
	}

	return nil
}

func (srv *Server) AddExpense(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := &AddExpense{}

	d := json.NewDecoder(r.Body)
	if err := d.Decode(req); err != nil {
		srv.writeResponseMessage(w, http.StatusBadRequest, err.Error())

		return
	}

	i, err := time.ParseDuration(req.Interval)
	if err != nil {
		srv.writeResponseMessage(w, http.StatusBadRequest, err.Error())

		return
	}

	re := &domain.AddExpenseReq{
		CreationDate: req.CreationDate,
		Label:        req.Label,
		Value:        req.Value,
		Interval:     i,
	}

	if err := srv.deps.App.AddExpense(r.Context(), re); err != nil {
		srv.writeResponseMessage(w, http.StatusInternalServerError, err.Error())

		return
	}

	srv.writeResponseMessage(w, http.StatusOK, "expense added")
}

func (srv *Server) GetAllExpenses(w http.ResponseWriter, r *http.Request) {
	e, err := srv.deps.App.GetAllExpenses(r.Context())
	if err != nil {
		srv.writeResponseMessage(w, http.StatusInternalServerError, err.Error())

		return
	}

	b, err := json.Marshal(e)
	if err != nil {
		srv.writeResponseMessage(w, http.StatusInternalServerError, err.Error())

		return
	}

	srv.writeResponseMessage(w, http.StatusOK, string(b))
}

func (srv *Server) writeResponseMessage(w http.ResponseWriter, statusCode int, message string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(message)) //nolint:errcheck
}
