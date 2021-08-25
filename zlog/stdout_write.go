package zlog

import "os"

type StdoutWriter struct {
	*AccessWriter
}

// NewStdoutWriter will write to stdout.
func NewStdoutWriter(options ...Option) *StdoutWriter {
	w := NewAccessWriter(func(w *AccessWriter){
		w.FileOut = os.Stderr
	})

	for _, opt := range options {
		opt(w)
	}

	return &StdoutWriter{w}
}

func (w *StdoutWriter) WriteLevel(l Level, p []byte) (n int, err error) {

	return w.Write(p)
}
