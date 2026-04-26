package webauthn

import (
	"strconv"

	"github.com/go-webauthn/webauthn/webauthn"
)

var _ webauthn.User = (*User)(nil)

// User implements webauthn.User interface.
type User struct {
	id         int
	name       string
	Credential []Credential
}

// Credential is a type alias for webauthn.Credential.
type Credential webauthn.Credential

// New creates User for WebAuthn.
func NewUser(id int, name string, c ...Credential) User {
	return User{
		id:         id,
		name:       name,
		Credential: c,
	}
}

// WebAuthnID provides user ID, Maximum 64 bytes.
func (u User) WebAuthnID() []byte {
	return []byte(strconv.Itoa(u.id))
}

// WebAuthnName provides name during registration.
func (u User) WebAuthnName() string {
	return u.name
}

// WebAuthnDisplayName provides display name during registration.
// Use name as display name temporarily.
func (u User) WebAuthnDisplayName() string {
	return u.name
}

// WebAuthnCredentials provides the list of Credential objects owned by the user.
func (u User) WebAuthnCredentials() []webauthn.Credential {
	var credentials []webauthn.Credential
	for _, c := range u.Credential {
		credentials = append(credentials, webauthn.Credential(c))
	}
	return credentials
}

// AddCredential adds Credential objects to the user.
func (u User) AddCredential(c ...Credential) User {
	u.Credential = append(u.Credential, c...)

	return u
}
