package intcode

// IntWriter is a writer of integers
type IntWriter interface {
	WriteInt(int)
}

// SimpleIntWriter is an IntWriter that remembers the values written to it.
// Not safe for usage for multiple goroutines.
type SimpleIntWriter struct {
	values []int
}

func NewSimpleIntWriter() SimpleIntWriter {
	return SimpleIntWriter{}
}

// WriteInt writes an value to this writer. Implements the IntWriter interface.
func (l *SimpleIntWriter) WriteInt(value int) {
	l.values = append(l.values, value)
}

// LastInt returns the last value written to this writer.
func (l *SimpleIntWriter) LastInt() int {
	if len(l.values) == 0 {
		return -1
	}

	return l.values[len(l.values)-1]
}

// LastInt returns the last value written to this writer.
func (l *SimpleIntWriter) FirstInt() int {
	if len(l.values) == 0 {
		return -1
	}

	return l.values[0]
}

// MaxInt returns the maximum integer written to this writer
func (l *SimpleIntWriter) MaxInt() int {
	if len(l.values) == 0 {
		return -1
	}

	max := l.values[0]

	for _, v := range l.values {
		if v > max {
			max = v
		}
	}

	return max
}

// Values returns a copy of the values written to this writer
func (l *SimpleIntWriter) Values() []int {
	res := make([]int, len(l.values))
	copy(res, l.values)
	return res
}

// MaxInt returns how many integers were written to this writer
func (l *SimpleIntWriter) Len() int {
	return len(l.values)
}

// DevNullIntWriter is an IntWriter that ignores every value written to it.
type DevNullIntWriter struct {
}

// WriteInt writes an value to this writer, which is ignored. Implements the IntWriter interface.
func (d *DevNullIntWriter) WriteInt(value int) {
}
