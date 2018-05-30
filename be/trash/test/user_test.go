package test

import(
	"testing"
	"encoding/json"

	"core_backend/gender"
	"core_backend/config"

	"gopkg.in/mgo.v2/bson"
)

var pathUser = map[string]string {
    "crud": "/user",
    "login": "/login",
    "userAccountTest": "/user/userAccountTest",
}

var preCreatedGender = PreCreateGender()

var validUser = map[string]string {
	"username": "useraccounttest",
    "email": "userAccountTest@gmail.com",
    "user": `{"username": "userAccountTest", "email": "userAccountTest@gmail.com", "password": "12345678", "gender": ` + preCreatedGender["gender"] + `}`,
    "login": `{"username": "userAccountTest", "password":"12345678"}`,
    "invalid-login": `{"username": "test", "password": "1234567"}`,
}

func PreCreateGender() map[string]string {
	_, _ = PostRequestString("/gender", HEADER_REQUEST_JSON, `{"description":"TESTE"}`)

	g := gender.Gender{}

	config.ConnectToDatabase()	

	config.FindOneByParameter(bson.M{"description": "TESTE"}, &g, "genders")

	b, _ := json.Marshal(g)

	config.RemoveFromDB(bson.M{"description": "TESTE"}, "genders")

	return map[string]string{
		"gender": string(b),	
	}
}

// OK TESTS
func TestCreateValidUser(t *testing.T) {

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, validUser["user"])

	assertEqual(t, resp, ExpectedResponses["created-user"])

	defer config.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}

func TestUserValidLogin(t *testing.T) {
	_, _ = PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, validUser["user"])

	resp, _ := PostRequest(pathUser["login"], HEADER_REQUEST_JSON, validUser["login"])
	response := ReaderToString(resp.Body)
	sessionHeader := resp.Header.Get("Authorization")

	assertEqual(t, response, ExpectedResponses["logged-in"])
	assertNotEqual(t, sessionHeader, "")

	defer config.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}

func TestUserValidUpdate(t *testing.T) {
    _, _ = PostRequest(pathUser["crud"], HEADER_REQUEST_JSON, validUser["user"])
    authRequest, _ := PostRequest(pathUser["login"], HEADER_REQUEST_JSON, validUser["login"])    

    request := `{ "email": "test@test.com" }`

    res, _ := PutRequestAuth(pathUser["userAccountTest"], request, authRequest.Header.Get("Authorization"))
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["updated"])

	defer config.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}

func TestUserValidDeactivate(t *testing.T) {
    _, _ = PostRequest(pathUser["crud"], HEADER_REQUEST_JSON, validUser["user"])
    authRequest, _ := PostRequest(pathUser["login"], HEADER_REQUEST_JSON, validUser["login"])    

    request := `{ "active": false }`

    res, _ := DeleteRequestAuth(pathUser["userAccountTest"], request, authRequest.Header.Get("Authorization"))
    response := ReaderToString(res.Body)
    assertEqualUserStruct(t, response, ExpectedResponses["destroyed"])

	defer config.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}

// BAD TESTS


func TestCreateUserBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badusername": "test", "bademail": "test", "password":"123456"`

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}

func TestCreateUserEmptyUsername(t *testing.T) {

	expected := `{"code":400,"message":"required:Username"}`
	request := `{"username": "", "email": "test@test.com", "password":"123456", "gender": ` + preCreatedGender["gender"] + `}`

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateUserEmptyEmail(t *testing.T) {

	expected := `{"code":400,"message":"required:Email"}`
	request := `{"username": "test", "email": "", "password":"123456", "gender": ` + preCreatedGender["gender"] + `}`

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateUserWrongEmail(t *testing.T) {

	expected := `{"code":400,"message":"email:Email"}`
	request := `{"username": "test", "email": "test", "password":"123456", "gender": ` + preCreatedGender["gender"] + `}`

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateUserWithoutGender(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"username": "test", "email": "test@test.com", "password":"12"}`

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateUserInvalidPasswordLength(t *testing.T) {

	expected := `{"code":400,"message":"password-length:Password"}`
	request := `{"username": "test", "email": "test@test.com", "password":"12", "gender": ` + preCreatedGender["gender"] + `}`

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateUsedUsername(t *testing.T) {
	expected := `{"code":400,"message":"used-username:Username"}`
	request := `{"username": "testusername1", "password":"123456", "email": "test123@aaaa.com", "gender": ` + preCreatedGender["gender"] + `}`
	requestD := `{"username": "testusername1", "password":"123456", "email": "test1234@aaaa.com", "gender": ` + preCreatedGender["gender"] + `}`
	_, _ = PostRequest(pathUser["crud"], HEADER_REQUEST_JSON, requestD)

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

	defer config.RemoveFromDB(bson.M{"username": "testusername1"}, "users")

}

func TestCreateUsedEmail(t *testing.T) {

	expected := `{"code":400,"message":"used-email:Email"}`
	request := `{"username": "testusername2", "password":"123456", "email": "test123@aaaa.com", "gender": ` + preCreatedGender["gender"] + `}`
	requestD := `{"username": "testusername2", "password":"123456", "email": "test123@aaaa.com", "gender": ` + preCreatedGender["gender"] + `}`
	_, _ = PostRequest(pathUser["crud"], HEADER_REQUEST_JSON, requestD)

	resp, _ := PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

	defer config.RemoveFromDB(bson.M{"username": "testusername2"}, "users")

}

func TestCreateUserInvalidLogin(t *testing.T) {

	_, _ = PostRequestString(pathUser["crud"], HEADER_REQUEST_JSON, validUser["user"])

	resp, _ := PostRequest(pathUser["login"], HEADER_REQUEST_JSON, validUser["invalid-login"])
	response := ReaderToString(resp.Body)

	assertEqual(t, response, ExpectedResponses["bad-login"])

	defer config.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")

}

func TestCreateUserValidUpdateWithWrongToken(t *testing.T) {
	_, _ = PostRequest(pathUser["crud"], HEADER_REQUEST_JSON, validUser["user"])

    request := `{ "email": "test@test.com" }`

    res, _ := PutRequestAuth(pathUser["userAccountTest"], request, "wrongtoken")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["unauthorized"])

	defer config.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")

}