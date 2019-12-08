package intcode

// IntWriter is a writer of integers
type IntWriter interface {
	WriteInt(int)
}

// SimpleIntWriter is an IntWriter that remembers the last value written to it.
type SimpleIntWriter struct {
	value int
}

// WriteInt writes an value to this writer. Implements the IntWriter interface.
func (l *SimpleIntWriter) WriteInt(value int) {
	l.value = value
}

// LastInt returns the last value written to this writer.
func (l *SimpleIntWriter) LastInt() int {
	return l.value
}
