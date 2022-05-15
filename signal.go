/*
Copyright (c) 2022 Kai Wells

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package exit

import "sync"

// A Signal is a synchronization mechanism to safely close a channel exactly
// once. Multiple goroutines can each call Send on the same Signal and listeners
// will only be notified of the first Send.
//
// Typically this would be used to notify a parent goroutine when any of its
// children have exited. Using a bare channel for this would be unsafe if more
// than one child tried to close the channel.
type Signal struct {
	m      sync.Mutex
	c      chan struct{}
	closed bool
}

// A NewSignal which is safe to Send multiple times from different goroutines.
// Listeners will only be notified of the first Send.
func NewSignal() *Signal {
	return &Signal{
		c: make(chan struct{}),
	}
}

// Send all listeners a notification exactly once.
//
// This is safe to call any number of times from multiple goroutines.
func (s *Signal) Send() {
	if s == nil {
		return
	}

	s.m.Lock()

	if !s.closed {
		s.closed = true
		close(s.c)
	}

	s.m.Unlock()
}

// Listen for a notification which will be sent exactly once.
// Panics if the Signal is nil.
//
// The channel will never have any values.
// The caller should only wait for this channel to close.
func (s *Signal) Listen() <-chan struct{} {
	if s == nil {
		panic("cannot listen to nil exit.Signal")
	}

	return s.c
}
