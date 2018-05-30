package test

import(
	"testing"
	"encoding/json"

	"core_backend/config"
	"core_backend/category"
	"core_backend/thread"
	"core_backend/gender"

	"gopkg.in/mgo.v2/bson"
)

var preCreate = PreCreateUserAndCategory()

var validThread = map[string]string {
    "title": "threadTest",
    "thread": `{"title": "threadTest", "content": "contentTest", "user": ` + preCreate["user"] + `, "category": ` + preCreate["category"] + `}`,
}

var pathThread = map[string]string {
    "crud": "/" + preCreate["categoryId"] + "/thread",
}

func PreCreateUserAndCategory() map[string]string {

	config.ConnectToDatabase()	

	_, _ = PostRequestString("/category", HEADER_REQUEST_JSON, `{"category": "categoryTest", "description": "descriptionTest"}`)

	c := category.Category{}

	config.FindOneByParameter(bson.M{"category": "categoryTest"}, &c, "categories")

	cJ, _ := json.Marshal(c)

	_, _ = PostRequestString("/gender", HEADER_REQUEST_JSON, `{"description": "genderTest"}`)

	g := gender.Gender{}

	config.FindOneByParameter(bson.M{"description": "genderTest"}, &g, "genders")

	gJ, _ := json.Marshal(g)

	config.RemoveFromDB(bson.M{"description": "genderTest"}, "genders")
	config.RemoveFromDB(bson.M{"category": "categoryTest"}, "categories")

	return map[string]string{
		"user": `{"username": "userAccountTest"}`,
		"category": string(cJ),
		"categoryId": c.Id.Hex(),
		"gender": string(gJ),
	}
}

// OK TESTS
func TestCreateValidThread(t *testing.T) {

	resp, _ := PostRequestString(pathThread["crud"], HEADER_REQUEST_JSON, validThread["thread"])

	assertEqual(t, resp, ExpectedResponses["created"])

	defer config.RemoveFromDB(bson.M{"title": validThread["title"]}, "threads")
}

func TestThreadValidUpdate(t *testing.T) {
    request := `{ "title": "titleTestUpdate" }`

    _, _ = PostRequestString(pathThread["crud"], HEADER_REQUEST_JSON, validThread["thread"])
	thread := thread.Thread{}
    config.FindOneByParameter(bson.M{"title": validThread["title"]}, &thread, "threads")
    pathThread["threadCrud"] = pathThread["crud"] + "/" + thread.Id.Hex()
    res, _ := PutRequestAuth(pathThread["threadCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["updated"])

	defer config.RemoveFromDB(bson.M{"title": "titleTestUpdate"}, "threads")
}

func TestThreadValidDeactivate(t *testing.T) {
    request := `{ "active": false }`

    _, _ = PostRequestString(pathThread["crud"], HEADER_REQUEST_JSON, validThread["thread"])
	thread := thread.Thread{}
    config.FindOneByParameter(bson.M{"title": validThread["title"]}, &thread, "threads")
    pathThread["threadCrud"] = pathThread["crud"] + "/" + thread.Id.Hex()
    res, _ := DeleteRequestAuth(pathThread["threadCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["destroyed"])

	defer config.RemoveFromDB(bson.M{"title": "threadTest"}, "threads")
}

// BAD TESTS
func TestCreateThreadBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badtitle": "test"`

	resp, _ := PostRequestString(pathThread["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreateThreadWrongBody(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badTitle": "test", "gender": "` + preCreate["gender"] + `"}`

	resp, _ := PostRequestString(pathThread["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}

func TestCreateThreadEmptyTitleField(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"title": "", "gender": "` + preCreate["gender"] + `"}`

	resp, _ := PostRequestString(pathThread["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}