package webauthn

import (
	"encoding/json"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/util/log"
)

var w *webauthn.WebAuthn

func init() {
	var err error
	config := config.Config.WebAuthn
	if w, err = webauthn.New(&webauthn.Config{
		RPID:          config.RPID,
		RPDisplayName: config.RPDisplayName,
		RPOrigins:     config.RPOrigins,
	}); err != nil {
		log.Errorln("New WebAuthn object error: ", err)
	}
}

// BeginRegister generates a new set of registration data for webauthn.
func BeginRegister(u User) ([]byte, []byte, error) {
	c, s, err := w.BeginRegistration(u, webauthn.WithAuthenticatorSelection(protocol.AuthenticatorSelection{
		UserVerification: protocol.VerificationRequired,
	}), webauthn.WithConveyancePreference(protocol.PreferNoAttestation))
	if err != nil {
		return nil, nil, err
	}
	creation, err := json.Marshal(c)
	if err != nil {
		return nil, nil, err
	}
	session, err := json.Marshal(s)
	if err != nil {
		return nil, nil, err
	}
	return creation, session, nil
}

// FinishRegister validates the registration data and returns the credential for webauthn.
func FinishRegister(u User, session, response []byte) (*webauthn.Credential, error) {
	var sessionData webauthn.SessionData
	if err := json.Unmarshal(session, &sessionData); err != nil {
		return nil, err
	}

	var ccr protocol.CredentialCreationResponse
	if err := json.Unmarshal(response, &ccr); err != nil {
		return nil, err
	}

	creationData, err := ccr.Parse()
	if err != nil {
		return nil, err
	}
	return w.CreateCredential(u, sessionData, creationData)
}

// BeginLogin generates a new set of login data for webauthn.
func BeginLogin(u User) ([]byte, []byte, error) {
	c, s, err := w.BeginLogin(u)
	if err != nil {
		return nil, nil, err
	}
	cred, err := json.Marshal(c)
	if err != nil {
		return nil, nil, err
	}
	session, err := json.Marshal(s)
	if err != nil {
		return nil, nil, err
	}
	return cred, session, nil
}

// FinishLogin validates the login data and returns the credential for webauthn.
func FinishLogin(u User, session, response []byte) (*webauthn.Credential, error) {
	var sessionData webauthn.SessionData
	if err := json.Unmarshal(session, &sessionData); err != nil {
		return nil, err
	}

	var car protocol.CredentialAssertionResponse
	if err := json.Unmarshal(response, &car); err != nil {
		return nil, err
	}

	creationData, err := car.Parse()
	if err != nil {
		return nil, err
	}
	return w.ValidateLogin(u, sessionData, creationData)
}
