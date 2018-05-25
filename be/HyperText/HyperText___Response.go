package HyperText

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"code"`
	Message    string `json:"message"`
}

type EnableStruct struct {
	Enable bool `json:"enable"`
}

var CustomResponses = map[string]string {
	//-------------------------------- UE --------------------------------//
	`success-Login`:          `Successfully Logged in`,
	`success-disabled`:       `Successfully disabled! The next time you log in you will be automatically reactivated`,
	`success-entityInsert`:   `Successfully Inserted`,
	`success-entityDelete`:   `Successfully Deleted`,
	`success-update`:         `Successfully Updated`,
	`wrong-usernameLength`:   `Username needs to contain 5 caracters or more`,
	`wrong-passwordLength`:   `Password needs to contain 6 caracters or more`,
	`wrong-verify`:           `Wrong username or password`,
	`wrong-usernameEmail`:    `Wrong Username and Email`,
	`wrong-usernamePassword`: `Wrong Username and Password`,
	`wrong-emailPassword`:    `Wrong Email an Password!"}`,
	`wrong-email`:            `Invalid Email`,
	`wrong-body`:             `Invalid Body`,
	`wrong-data`:             `Invalid Data`,
	`wrong-json`:             `Can't Decode Json`,
	`exist-username`:         `This username already exists`,
	`exist-email`:            `Email already Exist`,
	`error-enable`:           `Can't reactive your account now, try again later`,
	`error-tryLater`:         `try it again later`,
	`error-serverDB`:         `Several error on DB server`,
	`error-update`:           `Can't update your entity`,
	`error-disable`:          `Can't disable now! Try again later`,
	`notfound-entity`:        `This entity don't exist`,
	`empty-database`:         `Empty database`,
	`not-admin`:         			`You don't have Admin permission`,
}

var Responses = map[int]string {
	//1×× Informational
	100 : `Continue`,
	101 : `Switching Protocols`,
	102 : `Processing`,
	//2×× Success
	200 : `OK`,
	201 : `Created`,
	202 : `Accepted`,
	203 : `Non-authoritative Information`,
	204 : `No Content`,
	205 : `Reset Content`,
	206 : `Partial Content`,
	207 : `Multi-Status`,
	208 : `Already Reported`,
	226 : `IM Used`,
	//3×× Redirection
	300 : `Multiple Choices`,
	301 : `Moved Permanently`,
	302 : `Found`,
	303 : `See Other`,
	304 : `Not Modified`,
	305 : `Use Proxy`,
	307 : `Temporary Redirect`,
	308 : `Permanent Redirect`,
	//4×× Client Error
	400 : `Bad Request`,
	401 : `Unauthorized`,
	402 : `Payment Required`,
	403 : `Forbidden`,
	404 : `Not Found`,
	405 : `Method Not Allowed`,
	406 : `Not Acceptable`,
	407 : `Proxy Authentication Required`,
	408 : `Request Timeout`,
	409 : `Conflict`,
	410 : `Gone`,
	411 : `Length Required`,
	412 : `Precondition Failed`,
	413 : `Payload Too Large`,
	414 : `Request-URI Too Long`,
	415 : `Unsupported Media Type`,
	416 : `Requested Range Not Satisfiable`,
	417 : `Expectation Failed`,
	418 : `I'm a teapot`,
	421 : `Misdirected Request`,
	422 : `Unprocessable Entity`,
	423 : `Locked`,
	424 : `Failed Dependency`,
	426 : `Upgrade Required`,
	428 : `Precondition Required`,
	429 : `Too Many Requests`,
	431 : `Request Header Fields Too Large`,
	444 : `Connection Closed Without Response`,
	451 : `Unavailable For Legal Reasons`,
	499 : `Client Closed Request`,
	//5×× Server Error
	500 : `Internal Server Error`,
	501 : `Not Implemented`,
	502 : `Bad Gateway`,
	503 : `Service Unavailable`,
	504 : `Gateway Timeout`,
	505 : `HTTP Version Not Supported`,
	506 : `Variant Also Negotiates`,
	507 : `Insufficient Storage`,
	508 : `Loop Detected`,
	510 : `Not Extended`,
	511 : `Network Authentication Required`,
	599 : `Network Connect Timeout Error`,
}

func HttpResponse(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(payload)
}

func HttpErrorResponse(w http.ResponseWriter, code int, payload interface{}) {
	r := Response{
		StatusCode: code,
		Message:    payload.(string),
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(r)
}
