package server

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mchmarny/vulnasync/internal/handler"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	addressDefault = ":8080"
	logLevelEnvVar = "LOG_LEVEL"

	closeTimeout = 3
	readTimeout  = 10
	writeTimeout = 600
)

var (
	contextKey key
)

type key int

// Run starts the server and waits for termination signal.
func Run(name, version string) {
	initLogging(name, version)
	log.Info().Msg("starting server")

	mux := http.NewServeMux()

	// health check
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok")
	})

	// root handler
	mux.HandleFunc("/", handler.Process)

	address := addressDefault
	if val, ok := os.LookupEnv("PORT"); ok {
		address = fmt.Sprintf(":%s", val)
	}

	ctx := context.Background()

	run(ctx, mux, address)
}

// run starts the server and waits for termination signal.
func run(ctx context.Context, mux *http.ServeMux, address string) {
	server := &http.Server{
		Addr:              address,
		Handler:           mux,
		ReadHeaderTimeout: readTimeout * time.Second,
		WriteTimeout:      writeTimeout * time.Second,
		BaseContext: func(l net.Listener) context.Context {
			// adding server address to ctx handler functions receives
			return context.WithValue(ctx, contextKey, l.Addr().String())
		},
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("error listening for server")
		}
	}()
	log.Debug().Msg("server started")

	<-done
	log.Debug().Msg("server stopped")

	downCtx, cancel := context.WithTimeout(ctx, closeTimeout*time.Second)
	defer func() {
		cancel()
	}()

	if err := server.Shutdown(downCtx); err != nil {
		log.Fatal().Err(err).Msg("error shuting server down")
	}
}

func initLogging(name, version string) {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	lev, ok := os.LookupEnv(logLevelEnvVar)
	if ok {
		lvl, err := zerolog.ParseLevel(lev)
		if err != nil {
			log.Warn().Err(err).Msgf("error parsing log level: %s", lev)
		} else {
			zerolog.SetGlobalLevel(lvl)
		}
	}

	log.Logger = zerolog.New(os.Stdout).With().
		Str("name", name).
		Str("version", version).
		Logger()
}
