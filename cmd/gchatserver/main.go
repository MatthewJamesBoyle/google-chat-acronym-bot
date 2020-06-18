package main

import (
	"context"
	"fmt"
	"github.com/matthewjamesboyle/google-chat-acronym-bot/internal/logging"
	"github.com/matthewjamesboyle/google-chat-acronym-bot/internal/persist"
	"github.com/matthewjamesboyle/google-chat-acronym-bot/internal/respond"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	logger := logging.FromContext(ctx)

	if err := mainWithCtx(ctx); err != nil {
		logger.Fatal(err)
	}
}

func mainWithCtx(ctx context.Context) error {
	db := persist.NewInMemoryPersister()

	s, err := respond.NewService(db)
	if err != nil {
		return err
	}
	h := respond.HttpGchatRespondHandler{Svc: s}

	p := os.Getenv("PORT")
	if p == "" {
		p = ":80"
	}

	log.Println("starting Server")
	if err := http.ListenAndServe(fmt.Sprintf(":%s", p), h); err != nil {
		return err
	}
	return nil
}
