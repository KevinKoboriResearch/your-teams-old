package test

import(
	"testing"

	"core_backend/config"
	"core_backend/category"

	"gopkg.in/mgo.v2/bson"
)

// Caio, caso você esteja lendo isso, esteja ciente de que meus testes estão muito mal feitos porque não tive muito tempo para pensar em como melhorá-los

var pathCategory = map[string]string {
    "crud": "/category",
}

var validCategory = map[string]string {
    "name": "categoryTest",
    "category": `{"category": "categoryTest", "description": "descriptionTest"}`,
}

// OK TESTS
func TestCreateValidCategory(t *testing.T) {

	resp, _ := PostRequestString(pathCategory["crud"], HEADER_REQUEST_JSON, validCategory["category"])

	assertEqual(t, resp, ExpectedResponses["created"])

	defer config.RemoveFromDB(bson.M{"category": validCategory["name"]}, "categories")
}

func TestCategoryValidUpdate(t *testing.T) {
    request := `{ "category": "categoryTestUpdate" }`

    _, _ = PostRequestString(pathCategory["crud"], HEADER_REQUEST_JSON, validCategory["category"])
	c := category.Category{}
    config.FindOneByParameter(bson.M{"category": validCategory["name"]}, &c, "categories")
    pathCategory["categoryCrud"] = "/category/" + c.Id.Hex()
    res, _ := PutRequestAuth(pathCategory["categoryCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["updated"])

	defer config.RemoveFromDB(bson.M{"category": "categoryTestUpdate"}, "categories")
}

func TestCategoryValidDeactivate(t *testing.T) {
    request := `{ "active": false }`

    _, _ = PostRequestString(pathCategory["crud"], HEADER_REQUEST_JSON, validCategory["category"])
	c := category.Category{}
    config.FindOneByParameter(bson.M{"category": validCategory["name"]}, &c, "categories")
    pathCategory["categoryCrud"] = "/category/" + c.Id.Hex()
    res, _ := DeleteRequestAuth(pathCategory["categoryCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["destroyed"])

	defer config.RemoveFromDB(bson.M{"category": "categoryTest"}, "categories")
}

// BAD TESTS
func TestCreateCategoryBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badCategory": "test"`

	resp, _ := PostRequestString(pathCategory["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateCategoryWrongBody(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"badCategory": "test"}`

	resp, _ := PostRequestString(pathCategory["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}

func TestCreateCategoryEmptyCategoryField(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"category": ""}`

	resp, _ := PostRequestString(pathCategory["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}

func TestCreateCategoryEmptyDescriptionField(t *testing.T) {

	expected := `{"code":400,"message":"required:Description"}`
	request := `{"category": "test", "description": ""}`

	resp, _ := PostRequestString(pathCategory["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}