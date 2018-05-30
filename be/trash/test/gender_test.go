package test

import(
	"testing"

	"core_backend/config"
	"core_backend/gender"

	"gopkg.in/mgo.v2/bson"
)

var pathGender = map[string]string {
    "crud": "/gender",
}

var validGender = map[string]string {
    "description": "genderTest",
    "gender": `{"description": "genderTest"}`,
}

// OK TESTS
func TestCreateValidGender(t *testing.T) {

	resp, _ := PostRequestString(pathGender["crud"], HEADER_REQUEST_JSON, validGender["gender"])

	assertEqual(t, resp, ExpectedResponses["created"])

	defer config.RemoveFromDB(bson.M{"description": validGender["description"]}, "genders")
}

func TestGenderValidUpdate(t *testing.T) {
    request := `{ "description": "genderTestUpdate" }`

    _, _ = PostRequestString(pathGender["crud"], HEADER_REQUEST_JSON, validGender["gender"])
	g := gender.Gender{}
    config.FindOneByParameter(bson.M{"description": validGender["description"]}, &g, "genders")
    pathGender["genderCrud"] = "/gender/" + g.Id.Hex()
    res, _ := PutRequestAuth(pathGender["genderCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["updated"])

	defer config.RemoveFromDB(bson.M{"description": "genderTestUpdate"}, "genders")
}

func TestGenderValidDeactivate(t *testing.T) {
    request := `{ "active": false }`

    _, _ = PostRequestString(pathGender["crud"], HEADER_REQUEST_JSON, validGender["gender"])
	g := gender.Gender{}
    config.FindOneByParameter(bson.M{"description": validGender["description"]}, &g, "genders")
    pathGender["genderCrud"] = "/gender/" + g.Id.Hex()
    res, _ := DeleteRequestAuth(pathGender["genderCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["destroyed"])

	defer config.RemoveFromDB(bson.M{"description": "genderTest"}, "genders")
}

// BAD TESTS
func TestCreateGenderBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badDescription": "test"`

	resp, _ := PostRequestString(pathGender["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateGenderWrongBody(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"badDescription": "test"}`

	resp, _ := PostRequestString(pathGender["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}

func TestCreateGenderEmptyDescriptionField(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"description": ""}`

	resp, _ := PostRequestString(pathGender["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}