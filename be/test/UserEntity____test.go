package test

import(
	"testing"
	"encoding/json"
	"core_backend/Interface"
	"gopkg.in/mgo.v2/bson"
)

//ok
func TestPassSignUp(t *testing.T) {

	resp, _ := PostRequestString(pathUser["SignUp"], APPJASON_UTF_8, validUser["user"])

	assertEqual(t, resp, ExpectedResponses["success-insert"])

	defer Interface.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}

func TestUserValidLogin(t *testing.T) {
	_, _ = PostRequestString(pathUser["crud"], APPJASON_UTF_8, validUser["user"])

	resp, _ := PostRequest(pathUser["login"], APPJASON_UTF_8, validUser["login"])
	response := ReaderToString(resp.Body)
	sessionHeader := resp.Header.Get("Authorization")

	assertEqual(t, response, ExpectedResponses["logged-in"])
	assertNotEqual(t, sessionHeader, "")

	defer Interface.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}

func TestUserValidUpdate(t *testing.T) {
    _, _ = PostRequest(pathUser["crud"], APPJASON_UTF_8, validUser["user"])
    authRequest, _ := PostRequest(pathUser["login"], APPJASON_UTF_8, validUser["login"])

    request := `{ "email": "test@test.com" }`

    res, _ := PutRequestAuth(pathUser["userAccountTest"], request, authRequest.Header.Get("Authorization"))
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["updated"])

	defer Interface.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}

func TestUserValidDeactivate(t *testing.T) {
    _, _ = PostRequest(pathUser["crud"], APPJASON_UTF_8, validUser["user"])
    authRequest, _ := PostRequest(pathUser["login"], APPJASON_UTF_8, validUser["login"])

    request := `{ "active": false }`

    res, _ := DeleteRequestAuth(pathUser["userAccountTest"], request, authRequest.Header.Get("Authorization"))
    response := ReaderToString(res.Body)
    assertEqualUserStruct(t, response, ExpectedResponses["destroyed"])

	defer Interface.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")
}



//bad
func TestCreateUserBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badusername": "test", "bademail": "test", "password":"123456"`

	resp, _ := PostRequestString(pathUser["crud"], APPJASON_UTF_8, request)

	assertEqual(t, resp, expected)
}

func TestCreateUserEmptyUsername(t *testing.T) {

	expected := `{"code":400,"message":"required:Username"}`
	request := `{"username": "", "email": "test@test.com", "password":"123456"}`

	resp, _ := PostRequestString(pathUser["crud"], APPJASON_UTF_8, request)

	assertEqual(t, resp, expected)
}

func TestCreateUserEmptyEmail(t *testing.T) {

	expected := `{"code":400,"message":"required:Email"}`
	request := `{"username": "test", "email": "", "password":"123456"}`

	resp, _ := PostRequestString(pathUser["crud"], APPJASON_UTF_8, request)

	assertEqual(t, resp, expected)
}

func TestCreateUserWrongEmail(t *testing.T) {

	expected := `{"code":400,"message":"email:Email"}`
	request := `{"username": "test", "email": "test", "password":"123456"}`

	resp, _ := PostRequestString(pathUser["crud"], APPJASON_UTF_8, request)

	assertEqual(t, resp, expected)

}

func TestCreateUserInvalidPasswordLength(t *testing.T) {

	expected := `{"code":400,"message":"password-length:Password"}`
	request := `{"username": "test", "email": "test@test.com", "password":"12"}`

	resp, _ := PostRequestString(pathUser["crud"], APPJASON_UTF_8, request)

	assertEqual(t, resp, expected)

}

func TestCreateUsedUsername(t *testing.T) {
	expected := `{"code":400,"message":"used-username:Username"}`
	request := `{"username": "testusername1", "password":"123456", "email": "test123@aaaa.com"}`
	requestD := `{"username": "testusername1", "password":"123456", "email": "test1234@aaaa.com"}`
	_, _ = PostRequest(pathUser["crud"], APPJASON_UTF_8, requestD)

	resp, _ := PostRequestString(pathUser["crud"], APPJASON_UTF_8, request)

	assertEqual(t, resp, expected)

	defer Interface.RemoveFromDB(bson.M{"username": "testusername1"}, "users")

}

func TestCreateUsedEmail(t *testing.T) {

	expected := `{"code":400,"message":"used-email:Email"}`
	request := `{"username": "testusername2", "password":"123456", "email": "test123@aaaa.com"}`
	requestD := `{"username": "testusername2", "password":"123456", "email": "test123@aaaa.com"}`
	_, _ = PostRequest(pathUser["crud"], APPJASON_UTF_8, requestD)

	resp, _ := PostRequestString(pathUser["crud"], APPJASON_UTF_8, request)

	assertEqual(t, resp, expected)

	defer Interface.RemoveFromDB(bson.M{"username": "testusername2"}, "users")

}

func TestCreateUserInvalidLogin(t *testing.T) {

	_, _ = PostRequestString(pathUser["crud"], APPJASON_UTF_8, validUser["user"])

	resp, _ := PostRequest(pathUser["login"], APPJASON_UTF_8, validUser["invalid-login"])
	response := ReaderToString(resp.Body)

	assertEqual(t, response, ExpectedResponses["bad-login"])

	defer Interface.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")

}

func TestCreateUserValidUpdateWithWrongToken(t *testing.T) {
	_, _ = PostRequest(pathUser["crud"], APPJASON_UTF_8, validUser["user"])

    request := `{ "email": "test@test.com" }`

    res, _ := PutRequestAuth(pathUser["userAccountTest"], request, "wrongtoken")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["unauthorized"])

	defer Interface.RemoveFromDB(bson.M{"username": validUser["username"]}, "users")

}
