package web

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/kaiiorg/go-bif-examine/pkg/config"

	"github.com/rs/zerolog"
)

//go:embed static/*
var content embed.FS

type Web struct {
	config     *config.Config
	log        zerolog.Logger
	httpServer *http.Server
}

func New(conf *config.Config, log zerolog.Logger) *Web {
	web := &Web{
		config: conf,
		log:    log,
	}
	http.HandleFunc("/", web.handler)
	web.httpServer = &http.Server{
		Addr: fmt.Sprintf(":%d", web.config.Web.Port),
	}

	return web
}

func (web *Web) Run() {
	go func() {
		if err := web.httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			web.log.Fatal().Err(err).Msg("Web UI server exiting")
		}
	}()
}

func (web *Web) Close() {
	ctx, ctxCancel := context.WithTimeout(context.Background(), time.Second*3)
	web.httpServer.Shutdown(ctx)
	ctxCancel()
}

func (web *Web) handler(w http.ResponseWriter, r *http.Request) {
	var filePath string
	// Redirect special cases
	switch r.URL.Path {
	case "/":
		filePath = "static/index.html"
		web.log.Debug().Str("original", r.URL.Path).Str("new", filePath).Msg("Redirecting original request to new destination")
	default:
		filePath = fmt.Sprintf("static%s", r.URL.Path)
	}

	// Read the file from the embedded FS
	data, err := content.ReadFile(filePath)
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		web.log.Debug().Err(err).Str("filePath", filePath).Msg("Failed to find a requested file")
		return
	}

	// Set the Content-Type based on the file extension
	contentType := http.DetectContentType(data)
	w.Header().Set("Content-Type", contentType)

	// Write the file content to the response
	_, err = w.Write(data)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		web.log.Error().Err(err).Msg("Failed to write a response to a web request")
	}
}
