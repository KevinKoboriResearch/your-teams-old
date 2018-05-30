package test

import (
	"testing"
)

//____________________________ INSERT ________________________________________//
func TestUserEntityInsertWRONGBODY(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertWRONGBODY)
	response := respToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityInsertSUCCESS(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertSUCCESS)
	response := respToString(resp)
	compareResults(t, response, Success)
}

//____________________________ VERIFY ________________________________________//
func TestUserEntityVerifyWRONGBODY(t *testing.T) {
	resp := sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyWRONGBODY)
	response := respToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityVerifyWRONGPASSWORD(t *testing.T) {
	resp := sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyWRONGPASSWORD)
	response := respToString(resp)
	compareResults(t, response, errorDatabase)
}

func TestUserEntityVerifySUCCESS(t *testing.T) {
	resp := sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifySUCCESS)
	response := respToString(resp)
	compareResults(t, response, Success)
}

//____________________________ UPDATE SINGLE _________________________________//
func TestUserEntityUpdateSingleWRONGBODY(t *testing.T) {
	resp := sendRequestPut("http://localhost:8080/YourAccount/UpdateSingle/username=username", UserEntityVerifyWRONGBODY)
	response := respToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityUpdateSingleWRONGIMAGE(t *testing.T) {
	resp := sendRequestPut("http://localhost:8080/YourAccount/UpdateSingle/username=username", UserEntityUpdateSingleWRONGIMAGE)
	response := respToString(resp)
	compareResults(t, response, wrongValidation)
}

func TestUserEntityUpdateSingleSUCCESS(t *testing.T) {
	resp := sendRequestPut("http://localhost:8080/YourAccount/UpdateSingle/username=username", UserEntityVerifySUCCESS)
	response := respToString(resp)
	compareResults(t, response, Success)
}
