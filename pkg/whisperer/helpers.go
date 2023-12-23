package whisperer

import "time"

func (w *Whisperer) timeout(d time.Duration) {
	select {
	case <-w.ctx.Done():
	case <-time.After(d):
	}
}
