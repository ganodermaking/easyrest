package goish

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

// LogFormatterParams is the structure any formatter will be handed when time to log comes
type LogFormatterParams struct {
	Request *http.Request

	// TimeStamp shows the time after the server returns a response.
	TimeStamp time.Time
	// StatusCode is HTTP response code.
	StatusCode int
	// Latency is how much time the server cost to process a certain request.
	Latency time.Duration
	// ClientIP equals Context's ClientIP method.
	ClientIP string
	// Method is the HTTP method given to the request.
	Method string
	// Path is a path the client requests.
	Path string
	// ErrorMessage is set if error has occurred in processing the request.
	ErrorMessage string
}

// LogFormatter ...
type LogFormatter func(params LogFormatterParams) string

// LoggerConfig defines the config for Logger middleware.
type LoggerConfig struct {
	Formatter LogFormatter
	Output    io.Writer
}

// Logger ...
func Logger() HandlerFunc {
	return LoggerWithConfig(LoggerConfig{})
}

// LoggerWithConfig ...
func LoggerWithConfig(conf LoggerConfig) HandlerFunc {
	formatter := conf.Formatter
	out := conf.Output
	if out == nil {
		out = DefaultWriter
	}

	return func(c *Context) {
		// Start timer
		start := time.Now()
		path := c.Req.URL.Path
		raw := c.Req.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		param := LogFormatterParams{
			Request: c.Req,
		}

		// Stop timer
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)

		param.ClientIP = c.ClientIP()
		param.Method = c.Req.Method

		if raw != "" {
			path = path + "?" + raw
		}

		param.Path = path

		fmt.Fprint(out, formatter(param))
	}
}
