package util

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func CalculateSha256(reader io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func ConfigureLogging(logLevel, applicationName, applicationDescription string) {
	// Configure pretty logs
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stdout})

	zerologLevel, err := zerolog.ParseLevel(logLevel)
	if err != nil || zerologLevel == zerolog.NoLevel {
		zerologLevel = zerolog.InfoLevel
		log.Warn().Str("givenLogLevel", logLevel).Msg("Given an unexpected log level; defaulting to info level")
	}
	// Log application name and description just before changing the log level. This makes sure it always get printed
	log.Info().
		Str("applicationName", applicationName).
		Str("applicationDescription", applicationDescription).
		Send()

	zerolog.SetGlobalLevel(zerologLevel)
}

func WaitForInterrupt() {
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	sig := <-signalChan
	log.Warn().Str("signal", sig.String()).Msg("Received signal, exiting")
}
