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

func TestUserInsertEntitySuccess(t *testing.T) {
	expected := `{"code": 220, "message":"usuário inserido com sucesso!"}`
	username := `username`
	company := `@company.com`
	num := 1
	user := `{
						"username": "` + username + strconv.Itoa(num) + `",
						"email": "` + username + strconv.Itoa(num) + company + `",
						"password": "123456",
						"enable": true
				 	 }`

	r := bytes.NewReader([]byte(user))

	resp, _ := http.Post(myHttp.HTTP+myServer.SERVER_HOST+myRouter.SIGNUP, myHttp.APPJASON_UTF_8, r)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}

}
