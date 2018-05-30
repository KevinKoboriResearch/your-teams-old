package test

const (
	APPJASON       = "application/json"
	CHARSET_UTF_8  = "charset=UTF-8"
	APPJASON_UTF_8 = APPJASON + "; " + CHARSET_UTF_8
)

var (
		ExpectedResponses = map[string]string {
    "destroyed": `{"code":200,"message":"Successfully destroyed."}`,
    "updated": `{"code":200,"message":"Successfully updated."}`,
    "created": `{"code":200,"message":"Successfully created."}`,
    "bad-login": `{"code":401,"message":"Wrong username or password."}`,
    "unauthorized": `{"code":401,"message":"You're not authorized."}`,
    "created-user": `{"code":200,"message":"Account successfully created."}`,
    "logged-in": `{"code":200,"message":"Logged in."}`,
}

pathUserEntity = map[string]string {
		"SignUp": "/SignUp",
    "Login": "/Login",
    "UpdateSingle": "/YourAccount/UpdateSingle/username={username}",
		"Update": "/YourAccount/Update/username={username}",
		"Get": "/Search/User/username={username}",
		"GetAllWhile": "/Search/Users/position={position}&value={value}",
		"GetAllEnabled": "/Search/Users",
		"GetAll": "/Search/UsersAll",
		"DeleteUser": "/YourAccount/Delete/username={username}",



    "userAccountTest": "/user/userAccountTest",
}

	passUser = map[string]string {
	"username": "useraccounttest",
    "email": "userAccountTest@gmail.com",
    "user": `{"username": "userAccountTest", "email": "userAccountTest@gmail.com", "password": "12345678"}`,
    "login": `{"username": "userAccountTest", "password":"12345678"}`,
    "invalid-login": `{"username": "test", "password": "1234567"}`,
}
)
