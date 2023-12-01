package webauthn

import "github.com/go-webauthn/webauthn/webauthn"

// User implements webauthn.User interface.
type User struct {
	ID          []byte
	Name        string
	DisplayName string
	Credential  []Credential
}

// Credential is a type alias for webauthn.Credential.
type Credential webauthn.Credential

// New creates User for WebAuthn.
func NewUser(id, name, displayName string, c ...Credential) User {
	return User{
		ID:          []byte(id),
		Name:        name,
		DisplayName: displayName,
		Credential:  c,
	}
}

// WebAuthnID provides user ID, Maximum 64 bytes.
func (u User) WebAuthnID() []byte {
	return u.ID
}

// WebAuthnName provides name during registration.
func (u User) WebAuthnName() string {
	return u.Name
}

// WebAuthnDisplayName provides display name during registration.
func (u User) WebAuthnDisplayName() string {
	return u.DisplayName
}

// WebAuthnCredentials provides the list of Credential objects owned by the user.
func (u User) WebAuthnCredentials() []webauthn.Credential {
	var credentials []webauthn.Credential
	for _, c := range u.Credential {
		credentials = append(credentials, webauthn.Credential(c))
	}
	return credentials
}

// WebAuthnIcon is a deprecated option,
// this has been removed from the specification recommendation.
func (u User) WebAuthnIcon() string {
	return ""
}

// AddCredential adds Credential objects to the user.
func (u User) AddCredential(c Credential) User {
	u.Credential = append(u.Credential, c)
	return u
}
