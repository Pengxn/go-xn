package controller

import (
	"encoding/json"
	"errors"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/lib/cache"
	"github.com/Pengxn/go-xn/src/lib/webauthn"
	"github.com/Pengxn/go-xn/src/model"
)

// RegisterPage returns register html page.
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// BeginRegister returns register webauthn creation.
func BeginRegister(c *gin.Context) {
	username := c.PostForm("username")
	user := webauthn.NewUser(123, username)
	creation, session, err := webauthn.BeginRegister(user)
	if err != nil {
		slog.Error("BeginRegister error", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "BeginRegister error",
		})
		return
	}

	if err := cache.Add(username, session); err != nil {
		slog.Error("cache.Add error", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "cache.Add error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"creation": json.RawMessage(creation),
	})
}

// FinishRegisterRequest is request body for finish register.
type FinishRegisterRequest struct {
	Username   string          `json:"username"`
	Credential json.RawMessage `json:"credential"`
}

// FinishRegister validates register webauthn credential.
func FinishRegister(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		slog.Error("GetRawData error", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "GetRawData error",
		})
		return
	}

	var creation FinishRegisterRequest
	json.Unmarshal(data, &creation)

	user := webauthn.NewUser(123, creation.Username)

	session, exist := cache.Get(creation.Username)
	if !exist {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "session error",
		})
		return
	}

	credential, err := webauthn.FinishRegister(user, session.([]byte), creation.Credential)
	if err != nil {
		slog.Error("FinishRegister error", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "FinishRegister error",
		})
		return
	}

	cred := model.WebAuthnCredential{
		Name:            creation.Username,
		NickName:        creation.Username, // use name as nickname temporarily
		UserID:          123,
		CredentialID:    credential.ID,
		PublicKey:       credential.PublicKey,
		AttestationType: credential.AttestationType,
		AAGUID:          credential.Authenticator.AAGUID,
		SignCount:       credential.Authenticator.SignCount,
		CloneWarning:    credential.Authenticator.CloneWarning,
		Attachment:      string(credential.Authenticator.Attachment),
	}
	if !cred.Add() {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "Add data error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": credential,
	})
}

// LoginPage returns login html page.
func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{})
}

// BeginLogin returns login webauthn creation and session.
func BeginLogin(c *gin.Context) {
	username := c.PostForm("username")

	var (
		creation, session []byte
		err               error
	)
	if username == "" { // client-side discoverable login
		creation, session, err = webauthn.BeginDiscoverableLogin()
	} else {
		creation, session, err = beginLoginWithUsername(username)
	}
	if err != nil {
		slog.Error("BeginLogin error", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "BeginLogin error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":     http.StatusOK,
		"creation": json.RawMessage(creation),
		"session":  json.RawMessage(session),
	})
}

// beginLoginWithUsername returns login webauthn creation and session with username.
func beginLoginWithUsername(username string) ([]byte, []byte, error) {
	wc := model.WebAuthnCredential{Name: username}.Get()
	if len(wc) == 0 {
		return nil, nil, errors.New("webauthn credentials not found")
	}

	var credentials []webauthn.Credential
	for _, c := range wc {
		credentials = append(credentials, webauthn.Credential{
			ID:              c.CredentialID,
			PublicKey:       c.PublicKey,
			AttestationType: c.AttestationType,
			// TODO: add Authenticator fields
		})
	}

	user := webauthn.NewUser(123, username).AddCredential(credentials...)
	creation, session, err := webauthn.BeginLogin(user)
	if err != nil {
		return nil, nil, errors.New("BeginLogin error")
	}

	return creation, session, nil
}

// FinishRegisterRequest is request body for finish register.
type FinishLoginRequest struct {
	Username   string          `json:"username"`
	Credential json.RawMessage `json:"credential"`
	Session    json.RawMessage `json:"session"`
}

// FinishLogin validates login webauthn credential.
func FinishLogin(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		slog.Error("GetRawData error", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "GetRawData error",
		})
		return
	}

	var login FinishLoginRequest
	json.Unmarshal(data, &login)

	user := webauthn.NewUser(123, login.Username) // TODO: implement query user
	if login.Username == "" {
		_, err = webauthn.FinishLogin(user, login.Session, login.Credential)
	} else {
		// TODO: implement handler function for discoverable login
		_, err = webauthn.FinishDiscoverableLogin(nil, login.Session, login.Credential)
	}
	if err != nil {
		slog.Error("FinishLogin error", slog.Any("error", err))

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "FinishLogin error",
		})
		return
	}

	c.JSON(http.StatusNotAcceptable, gin.H{
		"code":    http.StatusOK,
		"message": "login success",
	})
}
