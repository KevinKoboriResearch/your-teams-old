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

func TestUserDeleteEntitySuccess(t *testing.T) {

	expected := `{"code": 240, "message":"usuário excluído com sucesso!"}`

	username := `username`

	num := 1

	req, _ := http.NewRequest("DELETE", myHttp.HTTP+myServer.SERVER_HOST+myRouter.DELETEUSERENTITY+"/"+username+strconv.Itoa(num), nil)

	resp, _ := http.DefaultClient.Do(req)

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	response := buf.String()

	if response != expected {
		t.Errorf("The http response was: %v Expected: %v", response, expected)
	}
}
