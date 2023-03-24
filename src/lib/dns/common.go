package dns

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"math/rand"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/Pengxn/go-xn/src/config"
	"github.com/Pengxn/go-xn/src/util/httplib"
)

type DNSCommon struct {
	URL             string
	Host            string
	action          string
	region          string
	timestamp       time.Time
	nonce           int
	signatureMethod string
}

// NewDNS news DNS instance with default settings.
func NewDNS() DNSCommon {
	return DNSCommon{
		URL:             "https://cns.api.qcloud.com/v2/index.php",
		Host:            "cns.api.qcloud.com",
		timestamp:       time.Now(),
		nonce:           rand.Intn(10000),
		signatureMethod: "HmacSHA256",
	}
}

// SetAction sets 'action' field in DNSCommon.
func (dns DNSCommon) SetAction(action string) DNSCommon {
	dns.action = action
	return dns
}

// do http request, https://cloud.tencent.com/document/product/302/7310
func (dns DNSCommon) do(params map[string]string) (*http.Response, error) {
	dnsConfig := config.Config.DNS
	// common params
	allParams := map[string]string{
		"Action":          dns.action,
		"Timestamp":       strconv.Itoa(int(dns.timestamp.Unix())),
		"Nonce":           strconv.Itoa(dns.nonce),
		"SecretId":        dnsConfig.SecretID,
		"SignatureMethod": dns.signatureMethod,
	}
	if dns.region != "" { // Optional param
		allParams["Region"] = dns.region
	}
	for k, v := range params { // other params
		allParams[k] = v
	}

	signatureRaw := fmt.Sprintf("POST%s/v2/index.php?%s", dns.Host, sortParams(allParams))
	signature := signHmacSha256(signatureRaw, dnsConfig.SecretKey)

	postParam := map[string][]string{}
	for k, v := range allParams { // convert map[string]string to map[string][]string
		postParam[k] = []string{v}
	}
	postParam["Signature"] = []string{signature}

	return httplib.New().SetParams(postParam).POST(dns.URL)
}

func sortParams(params map[string]string) string {
	var keys []string
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys) // map sort
	for i, v := range keys {
		keys[i] = v + "=" + params[v]
	}

	return strings.Join(keys, "&") // concat params by '&'
}

// signHmacSha256 returns HAMC SHA256
func signHmacSha256(message string, secret string) string {
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(message))

	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
