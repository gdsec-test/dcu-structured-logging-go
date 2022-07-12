package logger

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// HTTPInfo is the data structure that holds the required fields for
// http logging
type HTTPInfo struct {
	Scheme       string              `json:"scheme"`
	Duration     float64             `json:"duration"`
	RequestBody  string              `json:"request-body"`
	Status       int                 `json:"status"`
	Method       string              `json:"method"`
	Path         string              `json:"path"`
	Src          string              `json:"src"`
	Headers      map[string][]string `json:"headers"`
	ResponseBody string              `json:"response-body"`
}

// NewHTTPInfoLogger returns a zap logger object with required config
func NewHTTPInfoLogger() *zap.Logger {
	logger, _ := zap.Config{
		Encoding:    "json",
		Level:       zap.NewAtomicLevelAt(zapcore.InfoLevel),
		OutputPaths: []string{"stdout"},
	}.Build()
	return logger
}

// LogHTTPRequest is a helper function thats used to log an http request in middleware
func LogHTTPRequest(l *zap.Logger, r *http.Request, duration float64, responseStatusCode int, responseBody string) {
	bodyBytes, _ := ioutil.ReadAll(r.Body)
	bodyString := string(bodyBytes)
	r.Body.Close()
	httpInfo := HTTPInfo{
		Scheme:       "http",
		Duration:     duration,
		RequestBody:  bodyString,
		Status:       responseStatusCode,
		Method:       r.Method,
		Path:         r.RequestURI,
		Src:          r.RemoteAddr,
		Headers:      r.Header,
		ResponseBody: responseBody,
	}
	jsonHTTPInfo, _ := json.Marshal(httpInfo)
	l.Info("", zap.Any("http-info", json.RawMessage(jsonHTTPInfo)))
	l.Info(string(jsonHTTPInfo))
}
