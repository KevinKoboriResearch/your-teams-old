package test

import(
	"net/http"
	"bytes"
	"io"
	"strings"
	"os"
	"testing"
	"log"
	"be/Interface"
)

func TestMain(m *testing.M) {
    Interface.ConnectToDatabase()
	runTests := m.Run()
	os.Exit(runTests)
}

func PostRequest(path string, typeRequest string, entity string) (response *http.Response, err error) {
    return http.Post("http://" + Interface.SERVER_HOST + path, typeRequest, StringToReader(entity))
}

func PostRequestString(path string, typeRequest string, entity string) (response string, err error) {
    resp, err := http.Post("http://" + Interface.SERVER_HOST + path, typeRequest, StringToReader(entity))

    if err != nil {
        log.Print("[ERROR] not found: ", err)
        return
    }

    response = ReaderToString(resp.Body)

    return
}

func StringToReader(s string) *bytes.Reader {
    return bytes.NewReader([]byte(s))
}

func ReaderToString(r io.Reader) string {
    buf := new(bytes.Buffer)
    buf.ReadFrom(r)
    text := buf.String()
    return strings.TrimSuffix(text, "\n")
}

func assertEqual(t *testing.T, received interface{}, expected interface{}) {
    if received != expected {
        t.Errorf("Received: %v Expected: %v", received, expected)
    }
}

func assertNotEqual(t *testing.T, received interface{}, expected interface{}) {
    if received == expected {
        t.Errorf("Received: %v Expected: %v", received, expected)
    }
}

func assertEqualUserStruct(t *testing.T, received string, expected string) {
    if !strings.Contains(received, expected) {
        t.Errorf("Received: %v Expected: %v", received, expected)
    }
}

func PutRequestAuth(path string, request string, token string) (res *http.Response, err error) {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodPut, "http://" + Interface.SERVER_HOST + path, StringToReader(request))
    if token != "" {
        req.Header.Set("Authorization", token)
    }
    return client.Do(req)
}

func DeleteRequestAuth(path string, request string, token string) (res *http.Response, err error) {
    client := &http.Client{}
    req, _ := http.NewRequest(http.MethodDelete, "http://" + Interface.SERVER_HOST + path, StringToReader(request))
    req.Header.Set("Authorization", token)
    return client.Do(req)
}
