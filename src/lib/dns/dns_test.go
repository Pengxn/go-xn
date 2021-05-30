package dns

import (
	"net/http"
	"testing"

	. "github.com/agiledragon/gomonkey/v2"
	. "github.com/smartystreets/goconvey/convey"
)

func TestDomainList(t *testing.T) {
	doFunc := func(_ DNSCommon, _ map[string]string) (*http.Response, error) {
		return &http.Response{}, nil
	}
	patch := ApplyFunc(DNSCommon.do, doFunc)
	defer patch.Reset()
	patch.ApplyFunc(handle, func(_ *http.Response, _ error) (string, error) {
		return `{
"info": {"domain_total": 18}, "domains": [
{"id": 53980930, "status": "enable", "group_id": "1", "searchengine_push": "no",
"is_mark": "no", "ttl": "600", "cname_speedup": "disable", "remark": "",
"created_on": "2017-02-08 18:05:22", "updated_on": "2017-02-08 18:05:22",
"q_project_id": 0, "punycode": "yizero.wang", "ext_status": "dnserror",
"src_flag": "QCLOUD", "name": "yizero.wang", "grade": "DP_Free",
"grade_title": "新免费套餐", "is_vip": "no", "owner": "100000@qq.com",
"records": "2", "min_ttl": 600}
]}`, nil
	})

	data, err := DomainList(map[string]string{})

	Convey("Test DomainList.", t, func() {
		So(data, ShouldResemble, DomainListData{
			Info: struct {
				DomainTotal int `json:"domain_total"`
			}{DomainTotal: 18},
			Domains: []struct {
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
			}{{
				ID:               53980930,
				Status:           "enable",
				GroupID:          "1",
				SearchEnginePush: "no",
				IsMark:           "no",
				TTL:              "600",
				CnameSpeedup:     "disable",
				Remark:           "",
				CreatedOn:        "2017-02-08 18:05:22",
				UpdatedOn:        "2017-02-08 18:05:22",
				QProjectID:       0,
				PunyCode:         "yizero.wang",
				ExtStatus:        "dnserror",
				SrcFlag:          "QCLOUD",
				Name:             "yizero.wang",
				Grade:            "DP_Free",
				GradeTitle:       "新免费套餐",
				IsVip:            "no",
				Owner:            "100000@qq.com",
				Records:          "2",
				MinTTL:           600,
			}},
		})
		So(err, ShouldBeNil)
	})
}

func TestRecordList(t *testing.T) {
	doFunc := func(_ DNSCommon, _ map[string]string) (*http.Response, error) {
		return &http.Response{}, nil
	}
	patch := ApplyFunc(DNSCommon.do, doFunc)
	defer patch.Reset()
	patch.ApplyFunc(handle, func(_ *http.Response, _ error) (string, error) {
		return `{"domain": {
"id": "55309561", "name": "yizeroapitest.com", "punycode": "yizeroapitest.com",
"grade": "DP_Free", "owner": "0000@qq.com", "ext_status": "dnserror",
"ttl": 600, "min_ttl": 600, "dnspod_ns": ["f1g1ns1.dnspod.net", "f1g1ns2.dnspod.net"],
"status": "enable", "q_project_id": 0},
"info": {"sub_domains": "4", "record_total": "4"},
"records": [{"id": 281628246, "ttl": 86400, "value": "f1g1ns1.dnspod.net.", "enabled": 1,
"status": "enabled", "updated_on": "2017-02-20 10:15:47", "q_project_id": 0, "name": "@",
"line": "\u9ed8\u8ba4", "line_id": "0", "type": "NS", "remark": "", "mx": 0}]}`, nil
	})

	data, err := RecordList("xn--02f.com")

	Convey("Test DomainList.", t, func() {
		So(data, ShouldResemble, RecordListData{
			Domain: struct {
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
			}{
				ID:         "55309561",
				Name:       "yizeroapitest.com",
				Punycode:   "yizeroapitest.com",
				Grade:      "DP_Free",
				Owner:      "0000@qq.com",
				ExtStatus:  "dnserror",
				TTL:        600,
				MinTTL:     600,
				DnspodNS:   []string{"f1g1ns1.dnspod.net", "f1g1ns2.dnspod.net"},
				Status:     "enable",
				QProjectID: 0,
			},
			Info: struct {
				SubDomains  string `json:"sub_domains"`
				RecordTotal string `json:"record_total"`
			}{
				SubDomains:  "4",
				RecordTotal: "4",
			},
			Records: []struct {
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
			}{{
				ID:         281628246,
				TTL:        86400,
				Value:      "f1g1ns1.dnspod.net.",
				Enable:     1,
				Status:     "enabled",
				UpdatedOn:  "2017-02-20 10:15:47",
				QProjectID: 0,
				Name:       "@",
				Line:       "默认",
				LineID:     "0",
				Type:       "NS",
				Remark:     "",
				MX:         0,
			}},
		})
		So(err, ShouldBeNil)
	})
}

func TestRecordCreate(t *testing.T) {
	doFunc := func(_ DNSCommon, _ map[string]string) (*http.Response, error) {
		return &http.Response{}, nil
	}
	patch := ApplyFunc(DNSCommon.do, doFunc)
	defer patch.Reset()
	patch.ApplyFunc(handle, func(_ *http.Response, _ error) (string, error) {
		return `{"record": {
			"id": "282529423",
			"name": "test",
			"status": "enable",
			"weight": null
		}}`, nil
	})

	data, err := RecordCreate(RecordParam{
		Domain:     "test.com",
		SubDomain:  "test",
		RecordType: "TXT",
		RecordLine: "默认",
		Value:      "THIS_IS_TEST_TXT_RECORD",
	})

	Convey("Test RecordCreate.", t, func() {
		So(data, ShouldResemble, RecordData{
			Record: struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				Value  string `json:"value"`
				Status string `json:"status"`
				Weight int    `json:"weight"`
			}{
				ID:     "282529423",
				Name:   "test",
				Status: "enable",
				Weight: 0,
			},
		})
		So(err, ShouldBeNil)
	})
}

func TestRecordModify(t *testing.T) {
	doFunc := func(_ DNSCommon, _ map[string]string) (*http.Response, error) {
		return &http.Response{}, nil
	}
	patch := ApplyFunc(DNSCommon.do, doFunc)
	defer patch.Reset()
	patch.ApplyFunc(handle, func(_ *http.Response, _ error) (string, error) {
		return `{"record": {
			"id": "282529938",
			"name": "test",
			"value": "112.112.21.21",
			"status": "enable",
			"weight": null
		}}`, nil
	})

	data, err := RecordModify("822150916", RecordParam{
		Domain:     "test.com",
		SubDomain:  "test",
		RecordType: "TXT",
		RecordLine: "默认",
		Value:      "UPDATE",
	})

	Convey("Test RecordModify.", t, func() {
		So(data, ShouldResemble, RecordData{
			Record: struct {
				ID     string `json:"id"`
				Name   string `json:"name"`
				Value  string `json:"value"`
				Status string `json:"status"`
				Weight int    `json:"weight"`
			}{
				ID:     "282529938",
				Name:   "test",
				Value:  "112.112.21.21",
				Status: "enable",
				Weight: 0,
			},
		})
		So(err, ShouldBeNil)
	})
}

func TestRecordStatus(t *testing.T) {
	doFunc := func(_ DNSCommon, _ map[string]string) (*http.Response, error) {
		return &http.Response{}, nil
	}
	patch := ApplyFunc(DNSCommon.do, doFunc)
	defer patch.Reset()
	patch.ApplyFunc(handle, func(_ *http.Response, _ error) (string, error) {
		return "", nil
	})

	Convey("Test RecordStatus.", t, func() {
		So(RecordStatus("test.com", "409200473", "disable"), ShouldBeNil)
	})
}
