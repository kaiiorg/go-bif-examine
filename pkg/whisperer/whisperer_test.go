package whisperer

import (
	"context"

	"github.com/rs/zerolog/log"
)

func newWithMocks() *Whisperer {
	// No mocks yet, so don't call http client or grpc yet!
	w := &Whisperer{
		log: log.With().Logger(),
	}
	w.ctx, w.ctxCancel = context.WithCancel(context.Background())

	return w
}
