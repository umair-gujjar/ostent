package ostential
import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
	"bufio"
	"net/http"
)

type logger struct {
	production bool
	logger     *log.Logger
	handler    http.Handler
	loggedLOCK sync.Mutex
	logged     map[string]bool
}

func Logged(production bool, loglogger *log.Logger, handler http.Handler) http.Handler {
	return &logger{
		production: production,
		logger:     loglogger,
		handler:    handler,
	}
}
func LoggedFunc(production bool, logger *log.Logger) func(http.Handler) http.Handler {
	return func(handler http.Handler) http.Handler {
		return Logged(production, logger, handler)
	}
}

func (lg *logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	lw := &loggedResponseWriter{ResponseWriter: w}
	lg.handler.ServeHTTP(lw, r)

	if lg.production {
		lg.productionLog(start, *lw, r)
		return
	}
	lg.log(start, *lw, r)
}

func (lg *logger) productionLog(start time.Time, w loggedResponseWriter, r *http.Request) {
	if w.status != 200 && w.status != 304 { // && r.URL.Path != "/ws" {}
		lg.log(start, w, r)
		return
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	lg.loggedLOCK.Lock()
	if _, ok := lg.logged[host]; ok {
		lg.loggedLOCK.Unlock()
		return
	}
	lg.logged[host] = true
	lg.loggedLOCK.Unlock()

	lg.logger.Printf("%s\tRequested from %s; subsequent successful requests will not be logged\n", time.Now().Format("15:04:05"), host)
}

var ZEROTIME, _ = time.Parse("15:04:05", "00:00:00")

func(lg *logger) log(start time.Time, w loggedResponseWriter, r *http.Request) {
	diff := time.Since(start)
	since := ZEROTIME.Add(diff).Format("5.0000s")

	code := fmt.Sprintf("%d", w.status)
	if w.status != 200 && w.status != 304 {
		code = statusLine(w.status)
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		host = r.RemoteAddr
	}

	lg.logger.Printf("%s\t%s\t%s\t%v\t%s\t%s\n", start.Format("15:04:05"), host, since, code, r.Method, r.URL.Path)
}

type loggedResponseWriter struct {
	http.ResponseWriter
	http.Flusher // ?
	status int
}

func (w *loggedResponseWriter) Flush() {
	if f, ok := w.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

func (w *loggedResponseWriter) WriteHeader(s int) {
	w.ResponseWriter.WriteHeader(s)
	w.status = s
}

func (w *loggedResponseWriter) Write(b []byte) (int, error) {
	if w.status == 0 {
		w.WriteHeader(http.StatusOK)
	}
	return w.ResponseWriter.Write(b)
}

func (w *loggedResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := w.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, fmt.Errorf("the ResponseWriter doesn't support the Hijacker interface")
}