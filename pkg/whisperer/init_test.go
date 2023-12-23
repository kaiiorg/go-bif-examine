package whisperer

import (
	"os/exec"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestCheckWhisperAvailability_Found(t *testing.T) {
	// Arrange
	w := newWithMocks()

	// Act
	// Making a big assumption that whatever system is running our tests has ls available
	err := w.checkWhisperAvailability("ls")

	// Assert
	require.NoError(t, err)
}

func TestCheckWhisperAvailability_NotFound(t *testing.T) {
	// Arrange
	w := newWithMocks()

	// Act
	err := w.checkWhisperAvailability(uuid.NewString())

	// Assert
	require.ErrorIs(t, err, exec.ErrNotFound)
}
