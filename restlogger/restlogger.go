package restlogger

import (
	"net/http"
)

//Wrap responsewriter to log status code in middleware
type LoggingResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

//Overriding WriteHeader function
func (lrw *LoggingResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

//Create a new log response writer
func NewLoggingResponseWriter(w http.ResponseWriter) *LoggingResponseWriter {
	return &LoggingResponseWriter{w, http.StatusOK}
}

type Restlogger struct {
	RequestId   string  `json:"request_id"`
	Scheme      string  `json:"scheme"`
	Duration    float64 `json:"duration"`
	RequestBody string  `json:"request-body"`
	AccountName string  `json:"accountname"`
}

type HttpInfo struct {
	Status  string      `json:"status"`
	Src     string      `json:"src"`
	Method  string      `json:"method"`
	Path    string      `json:"path"`
	Headers http.Header `json:"headers"`
}
