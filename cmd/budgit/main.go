package main

import (
	"context"
	"fmt"
	"log"

	"github.com/apelletant/budgit/pkg/core"
	"github.com/apelletant/budgit/pkg/domain"
	"github.com/apelletant/budgit/pkg/repository/pgsql"
	"github.com/apelletant/budgit/pkg/transport/http"
	"golang.org/x/sync/errgroup"
)

type app struct {
	cfg *Config
}

func main() {
	RunApp(context.Background())
}

func RunApp(ctx context.Context) {
	app := &app{
		cfg: parseConf(),
	}

	if err := app.run(ctx); err != nil {
		log.Print(err)
	}
}

func (a *app) run(ctx context.Context) error {
	ctxWithCancel, cancel := context.WithCancel(ctx)
	defer cancel()

	_, err := setupDeps(a.cfg)
	if err != nil {
		// TODO log
		return fmt.Errorf("setupDeps: %w", err)
	}

	srv, err := a.createServer()
	if err != nil {
		// TODO Logs
		return fmt.Errorf("http.New: %w", err)
	}

	errG, errCtx := errgroup.WithContext(ctxWithCancel)
	errG.Go(func() error {
		if err := srv.Run(errCtx); err != nil {
			return fmt.Errorf("sh.Run: %w", err)
		}

		return nil
	})

	if err := errG.Wait(); err != nil {
		return err // nolint: wrapcheck
	}
	return nil
}

func (a *app) createServer() (*http.Server, error) {
	cfg := &http.Config{
		Port: a.cfg.Port,
	}

	store := a.createExencesStore()

	app := core.New(store)

	deps := &http.Dependencies{
		App: app,
	}

	wh, err := http.New(deps, cfg)
	if err != nil {
		return nil, fmt.Errorf("http.New: %w", err)
	}

	return wh, nil
}

func (a *app) createExencesStore() domain.Store {
	return pgsql.New()
}
