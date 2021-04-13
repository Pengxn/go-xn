package whois

import (
	"errors"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// GetWhois gets domain whois information.
// Supported TLD list: https://data.iana.org/TLD/tlds-alpha-by-domain.txt
func GetWhois(domain string) (string, error) {
	tld := domain[strings.LastIndex(domain, ".")+1:]
	whoisServer, err := getWhoisServer(tld)
	if err != nil {
		return "", err
	}

	whois, err := sendWhoisSocket(domain, whoisServer)
	if err != nil {
		return "", err
	}

	return whois, nil
}

func getWhoisServer(tld string) (string, error) {
	resp, err := http.Get("https://www.iana.org/domains/root/db/" + tld + ".html")
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Status code is " + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	ret := regexp.MustCompile("(?i:<b>WHOIS Server:</b>(.*?)\n)").FindStringSubmatch(string(bodyData))

	return strings.TrimSpace(ret[1]), nil
}

// sendWhoisSocket send whois request to whois server.
func sendWhoisSocket(domain, whoisServer string) (string, error) {
	result, _, err := NewSocketConn(whoisServer+":43", domain+"\r\n", 20)
	if err != nil {
		return "", err
	}

	return result, nil
}

//NewSocketConn news SOCKET connection and contents to server.
func NewSocketConn(server, content string, timeout int64) (string, int, error) {
	conn, err := net.Dial("tcp", server)
	if err != nil {
		return "", 0, err
	}
	defer conn.Close()

	if err := conn.SetReadDeadline(time.Now().Add(1 * time.Second)); err != nil {
		return "", 0, err
	}

	if _, err = conn.Write([]byte(content)); err != nil {
		return "", 0, err
	}

	buf := make([]byte, 1024)
	result, resLen := "", 0
	for {
		n, errs := conn.Read(buf)
		if errs != nil {
			if errs == io.EOF {
				err = nil
			} else {
				err = errs
			}
			break
		}
		if n == 0 {
			break
		}
		if n < 1024 {
			buf = buf[0:n]
		}
		result += string(buf)
		resLen += n
	}

	return result, resLen, err
}
