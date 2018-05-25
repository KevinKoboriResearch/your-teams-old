package auth

import (
	"net/http"

	"be/HyperText"
    "github.com/codegangsta/negroni"
)

var routes = HyperText.Routes{{}}

func CreateAuthRoutes() HyperText.Routes {
	routes := HyperText.Routes{

	}

	return routes
}


func SetAuthenticatedMiddleware(r func(http.ResponseWriter, *http.Request)) (n *negroni.Negroni) {
    n = negroni.New(negroni.HandlerFunc(ValidateToken), negroni.Wrap(http.HandlerFunc(r)))
    return
}
