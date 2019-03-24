package semaphore

// Semaphore is a channel-backed construct allowing N operations
// to happen concurrently.
type Semaphore struct {
	c chan struct{}
}

// New returns a pointer to a new instance of Semaphore.
func New(concurrency int) (s *Semaphore) {
	return &Semaphore{
		c: make(chan struct{}, concurrency),
	}
}

// Run executes a function, blocking if N operations are executing.
func (s *Semaphore) Run(f func()) {
	s.c <- struct{}{}

	go func() {
		defer func() {
			<-s.c
		}()

		f()
	}()
}

// Wait ensures that all N operations are completed, by filling the
// channel.  This will block until all N operations are complete.
func (s *Semaphore) Wait() {
	// Fill the channel with empty structs, ensuring that there are
	// no existing operations running, as these would need to finish
	// before the channel could be written to.
	for i := 0; i < cap(s.c); i++ {
		s.c <- struct{}{}
	}

	// Drain the channel, allowing the semaphore to be used again.
	for i := 0; i < cap(s.c); i++ {
		<-s.c
	}
}
