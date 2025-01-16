package turnstile

var apiURL = "https://challenges.cloudflare.com/turnstile/v0/siteverify"

// Turnstile is a struct that holds the secret and the Turnstile API URL.
type Turnstile struct {
	Secret       string
	TurnstileURL string
}

// Response is the response from the Turnstile API, refer to
// https://developers.cloudflare.com/turnstile/get-started/server-side-validation/
type Response struct {
	// Success indicates if the challenge was passed.
	Success bool `json:"success"`
	// ErrorCodes is a list of errors that occurred. (failure only)
	ErrorCodes []string `json:"error-codes"`
	// ChallengeTs is the ISO timestamp for the time the challenge was solved.
	ChallengeTs string `json:"challenge_ts"`
	// Hostname is the hostname for which the challenge was served.
	Hostname string `json:"hostname"`
	// action is the customer widget identifier passed to the widget on the client side.
	// This is used to differentiate widgets using the same sitekey in analytics.
	// Its integrity is protected by modifications from an attacker. It is recommended to
	// validate that the action matches an expected value.
	Action string `json:"action"`
	// CData is the customer data passed to the widget on the client side.
	// This can be used by the customer to convey state. It is integrity protected by
	// modifications from an attacker.
	CData string `json:"cdata"`
	// EphemeralID returns the Ephemeral ID in siteverify. (enterprise only)
	EphemeralID string `json:"ephemeral_id"`
}

var (
	// The secret parameter was not passed.
	ErrMissingInputSecret = "missing-input-secret"
	// The secret parameter was invalid, did not exist,
	// or is a testing secret key with a non-testing response.
	ErrInvalidInputSecret = "invalid-input-secret"
	// The response parameter (token) was not passed.
	ErrMissingInputResponse = "missing-input-response"
	// The response parameter (token) is invalid or has expired.
	// Most of the time, this means a fake token has been used.
	// If the error persists, contact customer support.
	ErrInvalidInputResponse = "invalid-input-response"
	// The request was rejected because it was malformed.
	ErrBadRequest = "bad-request"
	// The response parameter (token) has already been validated before.
	// This means that the token was issued five minutes ago and is no longer valid,
	// or it was already redeemed.
	ErrTimeoutOrDuplicate = "timeout-or-duplicate"
	// An internal error happened while validating the response.
	// The request can be retried.
	ErrInternalError = "internal-error"
)
