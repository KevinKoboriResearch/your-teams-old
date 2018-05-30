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

func TestUserDeleteMultEntitySuccess(t *testing.T) {

	expected := `{"code": 240, "message":"usuário excluído com sucesso!"}`

	errorserver := `{"code": 507, "message":"Several error on DB server"}`

	username := `username`

	num := 0

	req, _ := http.NewRequest("DELETE", myHttp.HTTP+myServer.SERVER_HOST+myRouter.DELETEUSERENTITY+"/"+username, nil)

	resp, _ := http.DefaultClient.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	for response == expected {

		req, _ := http.NewRequest("DELETE", myHttp.HTTP+myServer.SERVER_HOST+myRouter.DELETEUSERENTITY+"/"+username+strconv.Itoa(num), nil)

		resp, _ := http.DefaultClient.Do(req)

		buf := new(bytes.Buffer)
		buf.ReadFrom(resp.Body)
		response = buf.String()

		num = num + 1

	}

	if response != errorserver {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
