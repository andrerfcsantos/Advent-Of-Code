package intcode

// IntPipe represents a pipe of integers. A pipe allows integers to be written to it and be read later, possibly
// by different processes. Can be used by multiple goroutines at the same time. Implements both the IntWriter
// and IntReader interfaces. The typical use of a pipe will be to use it as an IntWriter by a process and
// as an IntReader by another.
type IntPipe struct {
	buffer chan int
}

// NewIntPipe creates and initializes a new IntPipe
func NewIntPipe() IntPipe {
	return IntPipe{
		buffer: make(chan int, 1024),
	}
}

// WriteInt writes an int to a pipe. Implements the IntWriter interface.
func (p *IntPipe) WriteInt(i int) {
	p.buffer <- i
}

// ReadInt reads an int from the pipe. Implements the IntReader interface
func (p *IntPipe) ReadInt() int {
	return <-p.buffer
}

// Close closes the pipe
func (p *IntPipe) Close() {
	close(p.buffer)
}
