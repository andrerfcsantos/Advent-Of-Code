package day07_2019

// MaxWriter is a IntWriter that only keeps the maximum value written to it
type MaxWriter struct {
	max int
}

// NewMaxWriter returns a new max writer
func NewMaxWriter() MaxWriter {
	return MaxWriter{}
}

// WriteInt writes an int to this writer. Implements the IntWriter interface.
func (mw *MaxWriter) WriteInt(value int) {
	if value > mw.max {
		mw.max = value
	}
}

// Max gets the maximum value written to this writer
func (mw *MaxWriter) Max() int {
	return mw.max
}
