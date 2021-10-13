package httplib

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"time"
)

// ContentType is used to indicate the media type of the resource.
type ContentType string

const (
	Raw  ContentType = "application/raw"  // Raw MIME type
	XML  ContentType = "application/xml"  // XML MIME type
	JSON ContentType = "application/json" // JSON MIME type
	TEXT ContentType = "application/text" // Text MIME type
	HTML ContentType = "application/html" // HTML MIME type
)

// Config 配置文件
type Config struct {
	Timeout            time.Duration     // Time out
	DisableKeepAlives  bool              // Disable keep alives
	InsecureSkipVerify bool              // Skip secure verify
	Encode             bool              // Encode URL for GET method
	IsEncodeForGo      bool              // To encode URL by golang lib
	Charset            string            // Charset (default UTF-8)
	Certificates       []tls.Certificate // SSL certificate
	Proxy              func(*http.Request) (*url.URL, error)
}

// NewHttpConfig Overwrite default http settings
func NewHttpConfig() *Config {
	return &Config{
		Timeout:            10 * time.Second,
		DisableKeepAlives:  true,
		InsecureSkipVerify: true,
		Encode:             true,
		IsEncodeForGo:      false,
		Charset:            "UTF-8",
		Certificates:       nil,
	}
}
