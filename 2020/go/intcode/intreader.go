package intcode

// IntReader is an input reader for ints
type IntReader interface {
	ReadInt() int
}

// SimpleIntReader implements an IntReader that will provide to readers defined values. The values this reader provides
// are set when calling NewSimpleIntReader.
type SimpleIntReader struct {
	pos    int
	buffer []int
}

// NewSimpleIntReader returns a new SimpleIntReader that will provide to reader the values passed as arguments
func NewSimpleIntReader(values ...int) SimpleIntReader {
	var res SimpleIntReader
	res.buffer = append(res.buffer, values...)
	return res
}

// ReadInt read a value from this reader. Implements the IntReader interface.
func (r *SimpleIntReader) ReadInt() int {
	if len(r.buffer) == 0 {
		return 0
	}
	value := r.buffer[r.pos]
	r.pos = (r.pos + 1) % len(r.buffer)
	return value
}
