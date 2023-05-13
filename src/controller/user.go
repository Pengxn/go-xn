package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/Pengxn/go-xn/src/lib/cache"
	"github.com/Pengxn/go-xn/src/lib/webauthn"
	"github.com/Pengxn/go-xn/src/model"
	"github.com/Pengxn/go-xn/src/util/log"
)

// RegisterPage returns register html page.
func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{})
}

// BeginRegister returns register webauthn creation.
func BeginRegister(c *gin.Context) {
	username := c.PostForm("username")
	displayName := c.PostForm("displayName")
	user := webauthn.NewUser("123", username, displayName)
	creation, session, err := webauthn.BeginRegister(user)
	if err != nil {
		log.Errorf("BeginRegister error: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "BeginRegister error",
		})
		return
	}

	if err := cache.Add(username, session); err != nil {
		log.Errorf("cache.Add error: %v", err)

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
	Username    string          `json:"username"`
	DisplayName string          `json:"displayName"`
	Credential  json.RawMessage `json:"credential"`
}

// FinishRegister validates register webauthn credential.
func FinishRegister(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		log.Errorf("GetRawData error: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "GetRawData error",
		})
		return
	}

	var creation FinishRegisterRequest
	json.Unmarshal(data, &creation)

	user := webauthn.NewUser("123", creation.Username, creation.DisplayName)

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
		log.Errorf("FinishRegister error: %v", err)

		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": "FinishRegister error",
		})
		return
	}

	cred := model.WebAuthnCredential{
		Name:            creation.Username,
		NickName:        creation.DisplayName,
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
