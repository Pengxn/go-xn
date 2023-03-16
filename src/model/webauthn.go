package model

import "time"

// WebAuthnSession represents the WebAuthn session data.
type WebAuthnSession struct {
	ID          int64      `json:"ID" xorm:"pk autoincr 'ID'"`
	UserID      int64      `json:"user_id" xorm:"INDEX 'user_id'"`
	Challenge   []byte     `json:"challenge" xorm:"INDEX VARBINARY(1024) 'challenge'"`
	CreatedTime *time.Time `json:"created_unix" xorm:"created 'created_time'"`
	UpdatedTime *time.Time `json:"updated_unix" xorm:"updated 'updated_time'"`
}

// TableName returns a better table name for WebAuthnSession.
func (ws WebAuthnSession) TableName() string {
	return "webauthn_session"
}

// Add adds a temporary WebAuthn session data.
func (ws WebAuthnSession) Add() {
	// TODO
}

func (ws WebAuthnSession) GetByChallenge(challenge []byte) {
	// TODO
}

// WebAuthnCredential represents the WebAuthn credential data.
type WebAuthnCredential struct {
	ID              int64      `json:"ID" xorm:"pk autoincr"`
	Name            string     `json:"name" xorm:"unique(s)"`
	NickName        string     `json:"nick_name" xorm:"varchar(25) notnull 'nick_name'"`
	UserID          int64      `json:"user_id" xorm:"INDEX unique(s)"`
	CredentialID    []byte     `json:"credential_id" xorm:"INDEX VARBINARY(1024)"`
	PublicKey       []byte     `json:"public_key"`
	AttestationType string     `json:"attestation_type"`
	AAGUID          []byte     `json:"aaguid"`
	SignCount       uint32     `json:"sign_count" xorm:"BIGINT"`
	CloneWarning    bool       `json:"clone_warning"`
	CreatedTime     *time.Time `json:"created_unix" xorm:"INDEX created"`
	UpdatedTime     *time.Time `json:"updated_unix" xorm:"updated"`
}

// TableName returns a better table name for WebAuthnCredential.
func (cred WebAuthnCredential) TableName() string {
	return "webauthn_credential"
}
