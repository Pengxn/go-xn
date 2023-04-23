package model

import (
	"time"

	"github.com/Pengxn/go-xn/src/util/log"
)

// WebAuthnCredential represents the WebAuthn credential data.
type WebAuthnCredential struct {
	ID              int64      `json:"ID" xorm:"pk autoincr 'ID'"`
	Name            string     `json:"name" xorm:"unique(s) 'name'"`
	NickName        string     `json:"nick_name" xorm:"varchar(25) notnull 'nick_name'"`
	UserID          int64      `json:"user_id" xorm:"INDEX unique(s) notnull 'user_id'"`
	CredentialID    []byte     `json:"credential_id" xorm:"INDEX VARBINARY(1024) notnull 'credential_id'"`
	PublicKey       []byte     `json:"public_key" xorm:"notnull 'public_key'"`
	AttestationType string     `json:"attestation_type" xorm:"'attestation_type'"`
	AAGUID          []byte     `json:"aaguid" xorm:"notnull 'aaguid'"`
	SignCount       uint32     `json:"sign_count" xorm:"BIGINT 'sign_count'"`
	CloneWarning    bool       `json:"clone_warning" xorm:"'clone_warning'"`
	Attachment      string     `json:"attachment" xorm:"'attachment'"`
	CreatedTime     *time.Time `json:"created_time,omitempty" xorm:"INDEX created 'created_time'"`
	UpdatedTime     *time.Time `json:"updated_time,omitempty" xorm:"updated 'updated_time'"`
}

// TableName returns a better table name for WebAuthnCredential.
func (cred WebAuthnCredential) TableName() string {
	return "webauthn_credential"
}

// Add adds a new WebAuthn credential to `webauthn_credential` table.
func (cred WebAuthnCredential) Add() bool {
	db := orm.NewSession()
	defer db.Close()

	affected, err := db.InsertOne(cred)
	if err != nil {
		log.Errorf("WebAuthn Credential add throw error: %+v, param: %+v", err, cred)
	}

	return affected > 0
}

// Get gets WebAuthn credential list by the given WebAuthnCredential condition.
func (cred WebAuthnCredential) Get() []WebAuthnCredential {
	db := orm.NewSession()
	defer db.Close()

	var credList []WebAuthnCredential
	if err := db.Find(&credList, &cred); err != nil {
		log.Errorf("WebAuthn Credential get throw error: %+v, param: %+v", err, cred)
	}

	return credList
}
