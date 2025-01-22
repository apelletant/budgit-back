package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/apelletant/budgit/pkg/domain"
	"golang.org/x/sync/errgroup"
)

var (
	ErrMissingServer = errors.New("server should be initialized")
)

type Config struct {
	Port int
}

func (cfg *Config) validate() error {
	return nil
}

type Dependencies struct {
	App domain.App
}

func (d *Dependencies) validate() error {
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

	/*c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3333", "https://stalker.tools.dev.k8s.coyote.local"},
		AllowedMethods:   []string{"GET", "UPDATE", "PUT", "POST", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Authorization", "X-CSRF-Token", "Cookie"},
		AllowCredentials: true,
		Debug:            true,
	})*/

	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(srv.HelloWorld))
	mux.Handle("POST /expence", http.HandlerFunc(srv.AddExpence))
	mux.Handle("GET /expences", http.HandlerFunc(srv.GetAllExpences))
	// mux.Handle("POST /sign-in", c.Handler(http.HandlerFunc(h.signInHandlerFunc)))

	srv.s = &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	return srv, nil
}

func (srv *Server) Run(ctx context.Context) error {

	errG, errCtx := errgroup.WithContext(ctx)

	errG.Go(func() error {
		// TODO proper logs
		fmt.Printf("listening on port %d\n", srv.cfg.Port)

		err := srv.s.ListenAndServe()
		if err != nil {
			return fmt.Errorf("h.srv.ListenAndServe: %w", err)
		}

		return nil
	})

	<-errCtx.Done()

	if err := srv.s.Close(); err != nil {
		//clog.Logger.Error("h.srv.Close", clog.Error(err))
	}

	if err := errG.Wait(); err != nil {
		return fmt.Errorf("errG.Wait: %w", err)
	}

	return nil
}

func (srv *Server) HelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

func (srv *Server) AddExpence(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	req := &AddExpence{}

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

	re := &domain.AddExpenceReq{
		CreationDate: req.CreationDate,
		Label:        req.Label,
		Value:        req.Value,
		Interval:     i,
	}

	if err := srv.deps.App.AddExpence(r.Context(), re); err != nil {
		srv.writeResponseMessage(w, http.StatusInternalServerError, err.Error())

		return
	}

	srv.writeResponseMessage(w, http.StatusOK, "expence added")
}

func (srv *Server) GetAllExpences(w http.ResponseWriter, r *http.Request) {
	e, err := srv.deps.App.GetAllExpences(r.Context())
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
	w.Write([]byte(message))
}
