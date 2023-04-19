package totp

import (
	"github.com/pquerna/otp/totp"
)

// Generate generates a new TOTP secret, and returns it.
func Generate(username string) (string, error) {
	otp, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "go-xn",
		AccountName: username,
	})
	if err != nil {
		return "", err
	}

	return otp.Secret(), nil
}

// Validate validates the TOTP code with the secret.
func Validate(secret, code string) bool {
	return totp.Validate(code, secret)
}
