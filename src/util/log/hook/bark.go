// Package hook provides a [logrus] hook for sending logs to [Bark], [Telegram Bot] and [io.Writer].
//
// Deprecated: this package is deprecated and will be removed in the next release.
// Use [github.com/Pengxn/go-xn/src/util/log] instead.
//
// [logrus]: https://github.com/sirupsen/logrus
// [Bark]: https://github.com/Finb/Bark
// [Telegram Bot]: https://core.telegram.org/bots/api
package hook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

// BarkHook is a hook for logrus logging library to send logs to Bark.
type BarkHook struct {
	token   string
	timeout time.Duration
	levels  []logrus.Level
}

// NewBarkHook creates new BarkHook instance with given token.
//
// Deprecated: this function is deprecated and will be removed in the next release.
// refer to [#429]
//
// [#429]: https://github.com/Pengxn/go-xn/pull/429
func NewBarkHook(token string) *BarkHook {
	return &BarkHook{
		token:   token,
		timeout: 5 * time.Second,
		levels: []logrus.Level{
			logrus.WarnLevel,
			logrus.ErrorLevel,
			logrus.FatalLevel,
			logrus.PanicLevel,
		},
	}
}

// SetLevels sets the levels which this hook will trigger.
func (bh *BarkHook) SetLevels(levels []logrus.Level) {
	bh.levels = levels
}

// Levels define on which log levels this hook would trigger.
func (bh *BarkHook) Levels() []logrus.Level {
	return bh.levels
}

// Fire sends the log entry to Bark.
func (bh *BarkHook) Fire(entry *logrus.Entry) error {
	go func() {
		msg := bh.format(entry)
		if err := bh.sendMessage(msg); err != nil {
			fmt.Fprintf(os.Stderr, "failed to bark hook: %v\n", err)
		}
	}()

	return nil
}

// barkMessage is the message format for Bark.
type barkMessage struct {
	Title string `json:"title"`
	Body  string `json:"body"`
	Group string `json:"group"`
	Icon  string `json:"icon"`
}

// sendMessage sends the message to Bark,
// refer to https://bark.day.app/#/tutorial
func (bh *BarkHook) sendMessage(msg barkMessage) error {
	encoded, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), bh.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("https://api.day.app/%s", bh.token),
		bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to send to bark: %w", err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("response status code is %d, data: %s", response.StatusCode, body)
	}

	// TODO: handle response data
	fmt.Fprintf(os.Stdout, "bark response: %s\n", body)

	return nil
}

// format formats the log entry to bark message text.
func (bh *BarkHook) format(entry *logrus.Entry) barkMessage {
	var levelWithEmoji string
	switch entry.Level {
	case logrus.WarnLevel:
		levelWithEmoji = "⚠️ Warning"
	case logrus.ErrorLevel:
		levelWithEmoji = "❌ Error"
	case logrus.FatalLevel:
		levelWithEmoji = "🔥 Fatal"
	case logrus.PanicLevel:
		levelWithEmoji = "🚨 Panic"
	}

	format := logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	serialized, _ := format.Format(entry)

	return barkMessage{
		Title: fmt.Sprintf("[FYJ] %s", levelWithEmoji),
		Body:  string(serialized),
		Group: "fyj",
		Icon:  "https://github.com/Pengxn.png",
	}
}
