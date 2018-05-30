package myTest

import (
	"bytes"
	"lfkk/be/myHttp"
	"lfkk/be/myRouter"
	"lfkk/be/myServer"
	"net/http"
	"testing"
)

func TestUserGetEntitySuccess(t *testing.T) {

	username := `username5`

	expected := `[{"username":"` + username + `","email":"username5@company.com","password":"123456","enable":true}]`

	resp, _ := http.Get(myHttp.HTTP + myServer.SERVER_HOST + myRouter.GETUSERENTITY + "/" + username)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
