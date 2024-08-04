package mymiddleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

// MyLogger wraps the zerolog.Logger
type MyLogger struct {
	zerolog.Logger
}

// Logger is the global logger instance
var Logger MyLogger

// NewLogger initializes and returns a new MyLogger instance
func NewLogger() MyLogger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

	output.FormatLevel = func(i interface{}) string {
		level := strings.ToUpper(fmt.Sprintf("%-6s", i))
		switch level {
		case "INFO  ":
			return fmt.Sprintf("\033[32m| %s|\033[0m", level) // Green
		case "WARN  ":
			return fmt.Sprintf("\033[33m| %s|\033[0m", level) // Yellow
		case "ERROR ":
			return fmt.Sprintf("\033[31m| %s|\033[0m", level) // Red
		default:
			return fmt.Sprintf("| %s|", level)
		}
	}
	output.FormatFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s:", i)
	}
	output.FormatFieldValue = func(i interface{}) string {
		return fmt.Sprintf("%s", i)
	}
	output.FormatErrFieldName = func(i interface{}) string {
		return fmt.Sprintf("%s: ", i)
	}

	zerologger := zerolog.New(output).With().Caller().Timestamp().Logger()
	Logger = MyLogger{zerologger}
	return Logger
}

// LogInfo returns a zerolog event for info level logging
func (l *MyLogger) LogInfo() *zerolog.Event {
	return l.Info()
}

// LogError returns a zerolog event for error level logging
func (l *MyLogger) LogError() *zerolog.Event {
	return l.Error()
}

// LoggingMiddleware logs the requests and responses for Chi
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Log the request
		Logger.LogInfo().Fields(map[string]interface{}{
			"method": r.Method,
			"uri":    r.URL.Path,
			"query":  r.URL.RawQuery,
		}).Msg("Request")

		// Capture the response
		rw := &responseWriter{w, http.StatusOK}
		next.ServeHTTP(rw, r)

		// Log the response if there was an error
		if rw.status >= 400 {
			Logger.LogError().Fields(map[string]interface{}{
				"status": rw.status,
			}).Msg("Response")
		}
	})
}

// responseWriter is a wrapper to capture the status code
type responseWriter struct {
	http.ResponseWriter
	status int
}

func (rw *responseWriter) WriteHeader(statusCode int) {
	rw.status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}
