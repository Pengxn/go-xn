package slogger

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// nr_endpoint is the API endpoint for [NewRelic].
//
// [NewRelic]: https://newrelic.com
const nr_endpoint = "https://log-api.newrelic.com/log/v1"

// NewRelicWriter is a writer that sends log messages to NewRelic.
type NewRelicWriter struct {
	key     string
	timeout time.Duration
}

// NewRelic creates a new NewRelicWriter with the given API key.
func NewRelic(key string) *NewRelicWriter {
	return &NewRelicWriter{
		key:     key,
		timeout: defaultTimeout,
	}
}

// compile-time check that NewRelicWriter implements the [io.Writer] interface.
var _ io.Writer = (*NewRelicWriter)(nil)

// Write sends the log message to [NewRelic], it implements the [io.Writer] interface.
// It is called by the [log/slog] package when a log message is written.
//
// [NewRelic]: https://newrelic.com
func (nr *NewRelicWriter) Write(msg []byte) (n int, err error) {

	// TODO: concurrent write by goroutine

	if err = nr.send(msg); err != nil {
		fmt.Fprintf(os.Stderr, "[newrelic] log error: %v, msg: %s", err, msg)
	}

	return len(msg), nil
}

// send sends the log message to NewRelic using its [Log HTTP API].
// It is called by the [Write] method.
//
// [Log HTTP API]: https://docs.newrelic.com/docs/logs/log-api/introduction-log-api/#endpoint
func (nr *NewRelicWriter) send(msg []byte) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), nr.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, nr_endpoint, bytes.NewBuffer(msg))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Api-Key", nr.key) // or "License-Key"

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusAccepted {
		return fmt.Errorf("response status code is %d, data: %s", response.StatusCode, body)
	}

	return nil
}
