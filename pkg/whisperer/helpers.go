package whisperer

import (
	"time"
)

func (w *Whisperer) timeout(d time.Duration) error {
	select {
	case <-w.ctx.Done():
		return w.ctx.Err()
	case <-time.After(d):
		return nil
	}
}
