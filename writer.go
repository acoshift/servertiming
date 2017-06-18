package servertiming

import (
	"bufio"
	"fmt"
	"net"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	t time.Time
}

func (w *responseWriter) WriteHeader(code int) {
	diff := time.Now().Sub(w.t) / time.Millisecond
	w.Header().Set("Server-Timing", fmt.Sprintf(`total=%d; "Total Response Time"`, diff))
	w.ResponseWriter.WriteHeader(code)
}

// Push implements Pusher interface
func (w *responseWriter) Push(target string, opts *http.PushOptions) error {
	if w, ok := w.ResponseWriter.(http.Pusher); ok {
		return w.Push(target, opts)
	}
	return http.ErrNotSupported
}

// Flush implements Flusher interface
func (w *responseWriter) Flush() {
	if w, ok := w.ResponseWriter.(http.Flusher); ok {
		w.Flush()
	}
}

// CloseNotify implements CloseNotifier interface
func (w *responseWriter) CloseNotify() <-chan bool {
	if w, ok := w.ResponseWriter.(http.CloseNotifier); ok {
		return w.CloseNotify()
	}
	return nil
}

// Hijack implements Hijacker interface
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if w, ok := w.ResponseWriter.(http.Hijacker); ok {
		return w.Hijack()
	}
	return nil, nil, http.ErrNotSupported
}
