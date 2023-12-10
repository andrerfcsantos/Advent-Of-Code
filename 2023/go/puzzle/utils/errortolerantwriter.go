package utils

import "io"

// ErrorTolerantWriter is an io.Writer that is a wrapper around a regular io.Writer that
// has its writes write to the underlying writer, until an error occurs, at which point,
// further writes will do nothing.
// To check if an error occurred in any of the writes and how many bytes were written up until
// the first error happened, call w.Error() and w.Bytes() respectively.
type ErrorTolerantWriter struct {
	err    error
	n      int
	writer io.Writer
}

// NewErrorTolerantWriter returns a new ErrorTolerantWriter with the given writer
// as the underlying writer.
func NewErrorTolerantWriter(w io.Writer) *ErrorTolerantWriter {
	return &ErrorTolerantWriter{
		writer: w,
	}
}

// Write writes to the underlying writer if none of the previous writes returned an error,
// does nothing otherwise. The returned error is always nil, to check if any error happened in any
// of the writes, call w.Error().
func (w *ErrorTolerantWriter) Write(p []byte) (int, error) {
	if w.err != nil {
		return 0, nil
	}

	n, err := w.writer.Write(p)
	w.n += n
	if err != nil {
		w.err = err
		return n, nil
	}
	return n, nil
}

// Error returns nil if no error occurred in any of the writes, otherwise returns the first error
// that happened.
func (w *ErrorTolerantWriter) Error() error {
	return w.err
}

// Bytes returns the number of bytes written to the underlying writer.
// If an error occurred in any of the writes, returns the bytes written until the first error
// happened.
func (w *ErrorTolerantWriter) Bytes() int {
	return w.n
}
