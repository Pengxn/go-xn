package slogger

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// bark_endpoint is the API endpoint for [Bark].
const bark_endpoint = "https://api.day.app/"

// BarkWriter is a writer that sends log messages to Bark.
type BarkWriter struct {
	token   string
	timeout time.Duration
}

// NewBark creates a new BarkWriter instance with the given token.
func NewBark(token string) *BarkWriter {
	return &BarkWriter{
		token:   token,
		timeout: 5 * time.Second, // default 5 seconds timeout
	}
}

// compile-time check that BarkWriter implements the [io.Writer] interface.
var _ io.Writer = (*BarkWriter)(nil)

// Write sends the log message to [Bark], it implements the [io.Writer] interface.
// It is called by the [log/slog] package when a log message is written.
//
// [Bark]: https://github.com/finb/bark
func (bw *BarkWriter) Write(msg []byte) (n int, err error) {
	raw, err := bw.format(msg)
	if err != nil {
		return
	}

	n, err = bw.send(raw)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[bark] log error: %v, msg: %s", err, msg)
	}

	return
}

// send sends the log message to Bark using its [API].
// It is called by the [Write] method.
//
// [API]: https://bark.day.app/#/tutorial
func (bw *BarkWriter) send(msg []byte) (n int, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), bw.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, bark_endpoint+bw.token, bytes.NewBuffer(msg))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return
	}
	if response.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("response status code is %d, data: %s", response.StatusCode, body)
	}

	return len(msg), nil
}

// barkMessage is the message format for Bark.
type barkMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`  // message content
	Group string `json:"group"` // group by this field
	Icon  string `json:"icon"`
}

// format formats the log to bark message text.
func (bh *BarkWriter) format(msg []byte) ([]byte, error) {
	return json.Marshal(barkMessage{
		Title: "[FYJ] 🔥 LOG",
		Body:  string(msg),
		Group: "fyj",
		Icon:  "https://github.com/Pengxn.png",
	})
}
