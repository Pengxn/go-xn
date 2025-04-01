package hook

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

// TelegramHook is a hook for logrus logging library to send logs to Telegram.
type TelegramHook struct {
	token   string
	chatID  any
	timeout time.Duration
	levels  []logrus.Level
}

// NewTelegramHook creates new hook instance with given token.
//
// Deprecated: this function is deprecated and will be removed in the next release.
// refer to [#429]
//
// [#429]: https://github.com/Pengxn/go-xn/pull/429
func NewTelegramHook(token string) *TelegramHook {
	return &TelegramHook{
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

// SetChatID sets `chat_id` field in JSON payload sent to Telegram API.
func (th *TelegramHook) SetChatID(chatID any) {
	th.chatID = chatID
}

// SetLevels sets the levels which this hook will trigger.
func (th *TelegramHook) SetLevels(levels []logrus.Level) {
	th.levels = levels
}

// Levels define on which log levels this hook would trigger.
func (th *TelegramHook) Levels() []logrus.Level {
	return th.levels
}

// Fire sends the log entry to Telegram.
func (th *TelegramHook) Fire(entry *logrus.Entry) error {
	go func() {
		msg := message{
			ChatID:    th.chatID,            // TODO: limit chat_id type to int64 or string
			Text:      th.format(entry, ""), // TODO: add spoiler for sensitive data
			ParseMode: "MarkdownV2",
		}
		if err := th.sendMessage(msg); err != nil {
			fmt.Fprintf(os.Stderr, "failed to telegram hook: %v\n", err)
		}
	}()

	return nil
}

// message is JSON payload representation sent to Telegram API.
type message struct {
	ChatID    any    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

// sendMessage sends the message to Telegram API,
// refer to https://core.telegram.org/bots/api#sendmessage
func (th *TelegramHook) sendMessage(msg message) error {
	encoded, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), th.timeout)
	defer cancel()

	request, err := http.NewRequestWithContext(ctx, http.MethodPost,
		fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", th.token),
		bytes.NewBuffer(encoded))
	if err != nil {
		return err
	}
	request.Header.Set("Content-Type", "application/json")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return fmt.Errorf("failed to sendMessage to telegram: %w", err)
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
	fmt.Fprintf(os.Stdout, "telegram response: %s\n", body)

	return nil
}

// format formats the log entry to telegram message text.
func (th *TelegramHook) format(entry *logrus.Entry, spoiler string) string {
	var levelWithEmoji string
	switch entry.Level {
	case logrus.WarnLevel:
		levelWithEmoji = "⚠️ *Warning*"
	case logrus.ErrorLevel:
		levelWithEmoji = "❌ *Error*"
	case logrus.FatalLevel:
		levelWithEmoji = "🔥 *Fatal*"
	case logrus.PanicLevel:
		levelWithEmoji = "🚨 *Panic*"
	}

	if spoiler != "" {
		spoiler = fmt.Sprintf("||%s||\n", escape(spoiler))
	}

	format := logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	}
	serialized, _ := format.Format(entry)

	var msg string
	if len(serialized) > 0 {
		msg = "```\n" + escape(string(serialized)) + "\n```"
	}

	return fmt.Sprintf("%s\n%s%s", levelWithEmoji, spoiler, msg)
}

var reservedCharacters = "_*[]()~`>#+-=|{}.!"

// escape escapes special characters with preceding '\' character,
// refer to https://core.telegram.org/bots/api#markdownv2-style
func escape(s string) string {
	var buf strings.Builder
	for _, c := range s {
		if strings.ContainsRune(reservedCharacters, c) {
			buf.WriteString(`\`)
		}
		buf.WriteRune(c)
	}

	return buf.String()
}
