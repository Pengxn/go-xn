package turnstile

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// New creates a new Turnstile struct with the secret set.
func New(secret string) *Turnstile {
	return &Turnstile{
		Secret:       secret,
		TurnstileURL: apiURL,
	}
}

// Verify verifies a "cf-turnstile-response" data field, with an optional remote IP set.
func (t *Turnstile) Verify(response, remoteip string) (*Response, error) {
	values := url.Values{"secret": {t.Secret}, "response": {response}}
	if remoteip != "" {
		values.Set("remoteip", remoteip)
	}
	resp, err := http.PostForm(t.TurnstileURL, values)
	if err != nil {
		return nil, fmt.Errorf("HTTP error: %w", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("HTTP read error: %w", err)
	}

	r := Response{}
	err = json.Unmarshal(body, &r)
	if err != nil {
		return nil, fmt.Errorf("JSON error: %w", err)
	}

	return &r, nil
}
