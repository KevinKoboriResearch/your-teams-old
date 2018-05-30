package test

import (
	"be/HyperText"
	"bytes"
	"net/http"
	"testing"
)

const (
	APPJASON                  = "application/json"
	CHARSET_UTF_8             = "charset=UTF-8"
	APPJASON_UTF_8            = APPJASON + "; " + CHARSET_UTF_8
	User_Entity_Collection    = "user_entity"
	UserEntityInsertWRONGBODY = `{
									"username":"username",
									"email":"usernamecompany.com",
									"password":"123456789",
									"image":"http://www.xxx.jpg",
									"admin":true,
									"enable":false
								}`
	UserEntityInsertSUCCESS = `{
									"username":"username",
									"email":"username@company.com",
									"password":"123456789",
									"image":"http://www.xxx.jpg",
									"admin":true,
									"enable":false
								}`
	UserEntityVerifyWRONGBODY = `{
									"username":"username"
									"password":"123456789"
								}`
	UserEntityVerifyWRONGPASSWORD = `{
									"username":"username",
									"password":"12345678910"
								}`
	UserEntityVerifySUCCESS = `{
									"username":"username",
									"password":"123456789"
								}`
	UserEntityUpdateSingleWRONGBODY = `{
									"password":"123456789",
									"position":"email"
									"value":"kevin@company.com"
								}`
	UserEntityUpdateSingleWRONGIMAGE = `{
									"password":"123456789",
									"position":"image",
									"value":"hp/wwwxxx.jpg"
								}`
	UserEntityUpdateSingleSUCCESS = `{
									"password":"123456789",
									"position":"image",
									"value":"http://www.xxx.jpg"
								}`
)

var (
	wrongBody       = HyperText.TestResponses["wrong-body"]
	wrongValidation = HyperText.TestResponses["wrong-validation"]
	errorDatabase   = HyperText.TestResponses["error-database"]
	Success         = HyperText.TestResponses["success"]
)

func sendPost(path string, typeReq string, entity string) (resp *http.Response) {
	r := bytes.NewReader([]byte(entity))
	resp, _ = http.Post(path, typeReq, r)
	return resp
}

func sendRequestPut(path string, entity string) (resp *http.Response) {
	r := bytes.NewReader([]byte(entity))
	resp, _ = http.NewRequest("PUT", path, r)
	return resp
}

func sendRequestDel(path string, entity string) (resp *http.Request) {
	r := bytes.NewReader([]byte(entity))
	req, _ := http.NewRequest("DELETE", path, r)
	resp, _ = http.DefaultClient.Do(req)
	return resp
}

func respToString(resp *http.Response) (response string) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response = buf.String()
	return
}

func compareResults(t *testing.T, response string, expected string) {
	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
