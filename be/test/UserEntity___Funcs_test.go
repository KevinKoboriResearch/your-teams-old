package test

import (
	//	"net/http"
	"testing"
)

//____________________________ INSERT ________________________________________//
func TestUserEntityInsertWRONGBODY(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertWRONGBODY)
	response := responseToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityInsertSUCCESS(t *testing.T) {
	resp := sendPost("http://localhost:8080/SignUp", APPJASON_UTF_8, UserEntityInsertSUCCESS)
	response := responseToString(resp)
	compareResults(t, response, Success)
}

//____________________________ VERIFY ________________________________________//
func TestUserEntityVerifyWRONGBODY(t *testing.T) {
	resp := sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyWRONGBODY)
	response := responseToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityVerifyWRONGPASSWORD(t *testing.T) {
	resp := sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifyWRONGPASSWORD)
	response := responseToString(resp)
	compareResults(t, response, errorDatabase)
}

func TestUserEntityVerifySUCCESS(t *testing.T) {
	auth = sendPost("http://localhost:8080/Login", APPJASON_UTF_8, UserEntityVerifySUCCESS)
	response := responseToString(auth)
	compareResults(t, response, Success)
}

//____________________________ UPDATE _________________________________//
func TestUserEntityUpdateWRONGBODY(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/YourAccount/Update/username=username", UserEntityUpdateWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityUpdateWRONGBODY(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/YourAccount/Update/username=username", UserEntityUpdateWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityUpdateWRONGBODY(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/YourAccount/Update/username=username", UserEntityUpdateWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, wrongBody)
}

func TestUserEntityUpdateSingleWRONGBODY(t *testing.T) {
	resp, _ := sendPut("http://localhost:8080/YourAccount/UpdateSingle/username=username", UserEntityUpdateSingleWRONGBODY, auth.Header.Get("Authorization"))
	response := responseToString(resp)
	compareResults(t, response, wrongBody)
}

/*
//____________________________ UPDATE SINGLE _________________________________//
func TestUserEntityUpdateSingleWRONGIMAGE(t *testing.T) {
	req := sendRequestPut("http://localhost:8080/YourAccount/UpdateSingle/username=username", UserEntityUpdateSingleWRONGIMAGE)
	response := requestToString(req)
	compareResults(t, response, wrongValidation)
}

func TestUserEntityUpdateSingleSUCCESS(t *testing.T) {
	req := sendRequestPut("http://localhost:8080/YourAccount/UpdateSingle/username=username", UserEntityVerifySUCCESS)
	response := requestToString(req)
	compareResults(t, response, Success)
}
*/ /*
func TestUserValidUpdate(t *testing.T) {
	_, _ = PostRequest(user_path["crud"], user_responses["user"])
	authRequest, _ := PostRequest(user_path["login"], user_responses["login"])
	request := `{"email":"useraccounttestupdate@gmail.com" }`
	res, _ := PutRequest(user_path["crud-user"], request, authRequest.Header.Get("Authorization"))
	response := ReaderToString(res.Body)
	assertEqual(t, response, ExpectedResponses["updated"])
}
func PutRequest(path string, request string, token string) (res *http.Response, err error) {
	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodPut, "http://"+config.SERVER_HOST+path, StringToReader(request))
	if token != "" {
		req.Header.Set("Authorization", token)
	}
	return client.Do(req)
}
*/
