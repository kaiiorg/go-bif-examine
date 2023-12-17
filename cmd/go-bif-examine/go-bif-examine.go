package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/kaiiorg/go-bif-examine/pkg/bif_examine"
	"github.com/kaiiorg/go-bif-examine/pkg/config"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	applicationName        = "go-bif-examine"
	applicationDescription = "Examines BIF files used by Bioware's Infinity Engine. See https://github.com/kaiiorg/go-bif-examine"
)

var (
	logLevel   = flag.String("log-level", "info", "Zerolog log level to use; trace, debug, info, warn, error, panic, etc")
	configPath = flag.String("config", "./configs/example.hcl", "path to HCL config file")
)

func main() {
	// Parse CLI flags, configure logging, and load config file
	flag.Parse()
	configureLogging()
	conf, err := config.LoadFromFile(*configPath)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("path", *configPath).
			Msg("Failed to load config file")
	}

	// Create and start BifExamine
	log.Info().Msg("Starting go-bif-examine; use ctrl+c to exit")
	be, err := bif_examine.New(conf)
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to create go-bif-examine")
	}
	err = be.Run()
	if err != nil {
		log.Fatal().
			Err(err).
			Msg("Failed to start go-bif-examine")
	}

	waitForInterrupt()

	// Close BifExamine
	err = be.Close()
	if err != nil {
		log.Warn().
			Err(err).
			Msg("Failed to cleanly close go-bif-examine")
	}
}

func configureLogging() {
	// Configure pretty logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	zerologLevel, err := zerolog.ParseLevel(*logLevel)
	if err != nil || zerologLevel == zerolog.NoLevel {
		zerologLevel = zerolog.InfoLevel
		log.Warn().Str("givenLogLevel", *logLevel).Msg("Given an unexpected log level; defaulting to info level")
	}
	// Log application name and description just before changing the log level. This makes sure it always get printed
	log.Info().
		Str("applicationName", applicationName).
		Str("applicationDescription", applicationDescription).
		Send()

	zerolog.SetGlobalLevel(zerologLevel)
}

func waitForInterrupt() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	log.Warn().Str("signal", sig.String()).Msg("Received signal, exiting")
}
