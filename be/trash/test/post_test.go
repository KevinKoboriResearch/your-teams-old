package test

import(
	"testing"
	"encoding/json"

	"core_backend/config"
	"core_backend/category"
	"core_backend/thread"

	"gopkg.in/mgo.v2/bson"
)

var preCreatePost = PreCreateUserCategoryAndThread()

var validPost = map[string]string {
    "content": "contentTest",
    "post": `{"content": "contentTest", "thread_id": "` + preCreatePost["threadId"] + `", "category_id": "` + preCreatePost["categoryId"] +  `" ,"user": ` + preCreatePost["user"] + `}`,
}

var pathPost = map[string]string {
    "crud": "/" + preCreatePost["categoryId"] + "/thread/" + preCreatePost["threadId"] + "/post",
}

func PreCreateUserCategoryAndThread() map[string]string {

	config.ConnectToDatabase()	

	_, _ = PostRequestString("/category", HEADER_REQUEST_JSON, `{"category": "categoryTest", "description": "descriptionTest"}`)

	c := category.Category{}

	config.FindOneByParameter(bson.M{"category": "categoryTest"}, &c, "categories")

	cJ, _ := json.Marshal(c)

	t := thread.Thread{}

	_, _ = PostRequestString("/" + c.Id.Hex() + "/thread", HEADER_REQUEST_JSON, `{"title": "threadTest", "content": "contentTest", "user": {"username": "userAccountTest"}, "category": ` + string(cJ) + `}`)

	config.FindOneByParameter(bson.M{"title": "threadTest"}, &t, "threads")

	config.RemoveFromDB(bson.M{"category": "categoryTest"}, "categories")
	config.RemoveFromDB(bson.M{"title": "threadTest"}, "threads")

	return map[string]string{
		"user": `{"username": "userAccountTest"}`,
		"categoryId": c.Id.Hex(),
		"threadId": t.Id.Hex(),
	}
}

// OK TESTS
func TestCreateValidPost(t *testing.T) {

	resp, _ := PostRequestString(pathPost["crud"], HEADER_REQUEST_JSON, validPost["post"])

	assertEqual(t, resp, ExpectedResponses["created"])

	defer config.RemoveFromDB(bson.M{"content": validPost["content"]}, "posts")
}

func TestPostValidUpdate(t *testing.T) {
    request := `{ "content": "contentTestUpdate" }`

    _, _ = PostRequestString(pathPost["crud"], HEADER_REQUEST_JSON, validPost["post"])
	post := thread.Post{}
    config.FindOneByParameter(bson.M{"content": validPost["content"]}, &post, "posts")
    pathPost["postCrud"] = pathPost["crud"] + "/" + post.Id.Hex()
    res, _ := PutRequestAuth(pathPost["postCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["updated"])

	defer config.RemoveFromDB(bson.M{"content": "contentTestUpdate"}, "posts")
}

func TestPostValidDeactivate(t *testing.T) {
    request := `{ "active": false }`

  	_, _ = PostRequestString(pathPost["crud"], HEADER_REQUEST_JSON, validPost["post"])
	post := thread.Post{}
    config.FindOneByParameter(bson.M{"content": validPost["content"]}, &post, "posts")
    pathPost["postCrud"] = pathPost["crud"] + "/" + post.Id.Hex()
    res, _ := DeleteRequestAuth(pathPost["postCrud"], request, "")
    response := ReaderToString(res.Body)
    assertEqual(t, response, ExpectedResponses["destroyed"])

	defer config.RemoveFromDB(bson.M{"content": "contentTest"}, "posts")
}

// BAD TESTS
func TestCreatePostBadJson(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badcontent": "test"`

	resp, _ := PostRequestString(pathPost["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)

}

func TestCreatePostWrongBody(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"badContent": "test", "user": "` + preCreate["user"] + `"}`

	resp, _ := PostRequestString(pathPost["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}

func TestCreatePostEmptyTitleField(t *testing.T) {

	expected := `{"code":400,"message":"Wrong JSON."}`
	request := `{"content": "", "user": "` + preCreate["user"] + `"}`

	resp, _ := PostRequestString(pathPost["crud"], HEADER_REQUEST_JSON, request)

	assertEqual(t, resp, expected)
}