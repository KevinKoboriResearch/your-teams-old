package myTest

import (
	"bytes"
	"lfkk/be/myHttp"
	"lfkk/be/myRouter"
	"lfkk/be/myServer"
	"net/http"
	"strconv"
	"testing"
)

func TestUserUpdateEntitySuccess(t *testing.T) {
	username := `username`
	company := `@company.com`
	update := "Update"
	num := 0

	expected := `{
				"username":"username0",
				"email":"username0@company.com",
				"password":"123456Update",
				"enable":true
			}`

	user := `{
				"username":"` + username + strconv.Itoa(num) + `",
				"email":"` + username + strconv.Itoa(num) + company + `",
				"password":"123456` + update + `",
				"enable":true
			}`

	r := bytes.NewReader([]byte(user))

	resp, _ := http.NewRequest("PUT", myHttp.HTTP+myServer.SERVER_HOST+myRouter.UPDATEUSERENTITY, r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}
