package dns

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/tidwall/gjson"
)

// DomainListData DomainList data struct
type DomainListData struct {
	Info struct {
		DomainTotal int `json:"domain_total"`
	} `json:"info"`
	Domains []struct {
		ID               int    `json:"id"`
		Status           string `json:"status"`
		GroupID          string `json:"group_id"`
		SearchEnginePush string `json:"searchengine_push"`
		IsMark           string `json:"is_mark"`
		TTL              string `json:"ttl"`
		CnameSpeedup     string `json:"cname_speedup"`
		Remark           string `json:"remark"`
		CreatedOn        string `json:"created_on"`
		UpdatedOn        string `json:"updated_on"`
		QProjectID       int    `json:"q_project_id"`
		PunyCode         string `json:"punycode"`
		ExtStatus        string `json:"ext_status"`
		SrcFlag          string `json:"src_flag"`
		Name             string `json:"name"`
		Grade            string `json:"grade"`
		GradeTitle       string `json:"grade_title"`
		IsVip            string `json:"is_vip"`
		Owner            string `json:"owner"`
		Records          string `json:"records"`
		MinTTL           int    `json:"min_ttl"`
	} `json:"domains"`
}

// DomainList returns all domains with their details.
func DomainList(params map[string]string) (DomainListData, error) {
	var data DomainListData
	content, err := handle(NewDNS().SetAction("DomainList").do(params))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal([]byte(content), &data)

	return data, err
}

// RecordListData RecordList data struct
type RecordListData struct {
	Domain struct {
		ID         string   `json:"id"`
		Name       string   `json:"name"`
		Punycode   string   `json:"punycode"`
		Grade      string   `json:"grade"`
		Owner      string   `json:"owner"`
		ExtStatus  string   `json:"ext_status"`
		TTL        int      `json:"ttl"`
		MinTTL     int      `json:"min_ttl"`
		DnspodNS   []string `json:"dnspod_ns"`
		Status     string   `json:"status"`
		QProjectID int      `json:"q_project_id"`
	} `json:"domain"`
	Info struct {
		SubDomains  string `json:"sub_domains"`
		RecordTotal string `json:"record_total"`
	} `json:"info"`
	Records []struct {
		ID         int    `json:"id"`
		TTL        int    `json:"ttl"`
		Value      string `json:"value"`
		Enable     int    `json:"enabled"`
		Status     string `json:"status"`
		UpdatedOn  string `json:"updated_on"`
		QProjectID int    `json:"q_project_id"`
		Name       string `json:"name"`
		Line       string `json:"line"`
		LineID     string `json:"line_id"`
		Type       string `json:"type"`
		Remark     string `json:"remark"`
		MX         int    `json:"mx"`
	} `json:"records"`
}

// RecordList returns all DNS records of the domain.
func RecordList(domain string) (RecordListData, error) {
	var data RecordListData
	content, err := handle(NewDNS().SetAction("RecordList").do(map[string]string{
		"domain": domain,
	}))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal([]byte(content), &data)

	return data, err
}

// RecordData RecordCreate/RecordStatus/RecordModify data struct
type RecordData struct {
	Record struct {
		ID     string `json:"id"`
		Name   string `json:"name"`
		Value  string `json:"value"`
		Status string `json:"status"`
		Weight int    `json:"weight"`
	} `json:"record"`
}

// RecordCreate creates a new DNS record for domain.
func RecordCreate(domain string, params map[string]string) (string, error) {
	if domain == "" {
		return "", errors.New("")
	}
	params["domain"] = domain

	return handle(NewDNS().SetAction("RecordCreate").do(params))
}

// RecordStatus updates the status of DNS record.
func RecordStatus(domain string, recordID int, status string) (string, error) {
	params := map[string]string{
		"domain":   domain,
		"recordId": strconv.Itoa(recordID),
		"status":   status,
	}

	return handle(NewDNS().SetAction("RecordStatus").do(params))
}

// RecordModify updates DNS record for domain.
func RecordModify() (string, error) {
	params := map[string]string{
		"recordId": "",
		"status":   "",
	}

	return handle(NewDNS().SetAction("RecordModify").do(params))
}

// RecordDelete deletes DNS record.
func RecordDelete(domain string, recordID int) error {
	params := map[string]string{
		"domain":   domain,
		"recordId": strconv.Itoa(recordID),
	}
	_, err := handle(NewDNS().SetAction("RecordDelete").do(params))

	return err
}

// handle handles http response, and returns content and error.
func handle(resp *http.Response, err error) (string, error) {
	if err != nil {
		return "", err
	}
	if resp.StatusCode != http.StatusOK {
		return "", errors.New("Status code isn't 200")
	}
	defer resp.Body.Close()
	bodyData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	data, err := handleResponse(bodyData)
	if err != nil {
		return "", err
	}

	return data, nil
}

// handleResponse handles http content returned by http according to
// Tencent Cloud the specification.
func handleResponse(content []byte) (string, error) {
	j := gjson.ParseBytes(content)
	if j.Get("code").Int() != 0 {
		return "", errors.New(j.Get("message").String())
	}
	return j.Get("data").String(), nil
}

// ToJSON returns the map that marshals from the body bytes as json in response.
// It calls Response inner.
func ToJSON(data string, v interface{}) error {
	return json.Unmarshal([]byte(data), v)
}
