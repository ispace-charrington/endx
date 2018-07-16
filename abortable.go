package endx

import "sync"

// Abortable is a convenience type for aborting groups of goroutines
// at once. context.Context would probably be a better solution.
type Abortable struct {
	m  sync.Mutex
	ok bool
	c  chan struct{}
}

func (a *Abortable) ready() {
	if !a.ok {
		a.c = make(chan struct{})
	}
}

// Wait returns a channel that will be closed when a.Abort is called.
func (a *Abortable) Wait() chan struct{} {
	a.m.Lock()
	defer a.m.Unlock()
	a.ready()
	return a.c
}

// Abort closes the channel returned by a.Wait.
func (a *Abortable) Abort() {
	a.m.Lock()
	defer a.m.Unlock()
	a.ready()
	if a.c == nil {
		return
	}
	close(a.c)
	a.c = nil
}
