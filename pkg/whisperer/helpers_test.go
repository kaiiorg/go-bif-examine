package whisperer

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTimeout_ExitOnCtxCancel(t *testing.T) {
	// Arrange
	w := newWithMocks()
	w.ctxCancel()
	testCtx, testCtxCancel := context.WithTimeout(context.Background(), time.Second)

	// Act
	var err error
	go func() {
		err = w.timeout(5 * time.Second)
		testCtxCancel()
	}()
	<-testCtx.Done()

	// Assert
	require.ErrorIs(t, testCtx.Err(), context.Canceled)
	require.ErrorIs(t, err, context.Canceled)
}

func TestTimeout_ExitOnTimeout(t *testing.T) {
	// Arrange
	w := newWithMocks()
	testCtx, testCtxCancel := context.WithTimeout(context.Background(), time.Second)

	// Act
	var err error
	go func() {
		err = w.timeout(time.Millisecond)
		testCtxCancel()
	}()
	<-testCtx.Done()

	// Assert
	require.ErrorIs(t, testCtx.Err(), context.Canceled)
	require.NoError(t, err)
}
