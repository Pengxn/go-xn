package httplib

import (
	"crypto/tls"
	"net/http"
	"net/url"
	"os"
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

// Config is used to override the default http client settings.
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

// NewHttpConfig returns default http config.
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

// getProxyByEnv returns proxy settings by environment variables.
func getProxyByEnv() func(*http.Request) (*url.URL, error) {
	proxyEnvs := []string{
		"HTTP_PROXY", "http_proxy",
		"HTTPS_PROXY", "https_proxy",
		"SOCKS_PROXY", "socks_proxy",
		"ALL_PROXY", "all_proxy",
	}

	var proxy string
	for _, env := range proxyEnvs {
		if proxy != "" {
			break
		}
		proxy = os.Getenv(env)
	}

	if proxy == "" {
		return nil
	}

	return func(req *http.Request) (*url.URL, error) {
		return url.Parse(proxy)
	}
}
