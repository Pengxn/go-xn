package slogger

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
)

// TelegramWriter is a writer that sends log messages to Telegram by using Bot API.
type TelegramWriter struct {
	token   string
	chatID  any
	timeout time.Duration
}

// NewTelegram creates new TelegramWriter instance with the given token.
func NewTelegram(token string) *TelegramWriter {
	return &TelegramWriter{
		token:   token,
		timeout: 5 * time.Second, // default 5 seconds timeout
	}
}

// SetChatID sets `chat_id` field in JSON payload sent to Telegram API.
func (th *TelegramWriter) SetChatID(chatID any) {
	th.chatID = chatID
}

// compile-time check that TelegramWriter implements the [io.Writer] interface.
var _ io.Writer = (*TelegramWriter)(nil)

// Write sends the log message to [Telegram Bot API], it implements the [io.Writer] interface.
// It is called by the [log/slog] package when a log message is written.
//
// [Telegram Bot API]: https://core.telegram.org/bots/api
func (th *TelegramWriter) Write(msg []byte) (n int, err error) {
	m := message{
		ChatID:    th.chatID,          // TODO: limit chat_id type to int64 or string
		Text:      th.format(msg, ""), // TODO: add spoiler for sensitive data
		ParseMode: "MarkdownV2",
	}
	if err := th.sendMessage(m); err != nil {
		fmt.Fprintf(os.Stderr, "failed to telegram hook: %v\n", err)
	}

	return
}

// message is JSON payload representation sent to Telegram API.
type message struct {
	ChatID    any    `json:"chat_id"`
	Text      string `json:"text"`
	ParseMode string `json:"parse_mode,omitempty"`
}

// sendMessage sends the message to Telegram API,
// refer to https://core.telegram.org/bots/api#sendmessage
func (th *TelegramWriter) sendMessage(msg message) error {
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
func (th *TelegramWriter) format(raw []byte, spoiler string) string {
	levelWithEmoji := "[FYJ] 🔥 Log"

	serialized := raw

	var msg string
	if len(serialized) > 0 {
		msg = "```\n" + escape(string(serialized)) + "\n```"
	}

	if spoiler != "" {
		spoiler = fmt.Sprintf("||%s||\n", escape(spoiler))
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
