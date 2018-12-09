package httpsvr

import (
	"net/http"
)

//ServeHTTP servers http on given address
func ServeHTTP(addr string) error {

	return http.ListenAndServe(addr, nil)
}
