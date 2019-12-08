package day07

// MaxIntPipe represents a pipe of integers that also remembers the last value written to it
type MaxIntPipe struct {
	buffer chan int
	max    int
}

// NewMaxIntPipe creates and initializes a new MaxIntPipe
func NewMaxIntPipe() MaxIntPipe {
	return MaxIntPipe{
		buffer: make(chan int, 1024),
	}
}

// WriteInt writes an int to the pipe. Implements the IntWriter interface.
func (p *MaxIntPipe) WriteInt(i int) {
	if i > p.max {
		p.max = i
	}
	p.buffer <- i
}

// ReadInt reads an int from the pipe. Implements the IntReader interface
func (p *MaxIntPipe) ReadInt() int {
	return <-p.buffer
}

// Max returns the max value written on this pipe
func (p *MaxIntPipe) Max() int {
	return p.max
}

// Close closes the pipe
func (p *MaxIntPipe) Close() {
	close(p.buffer)
}
