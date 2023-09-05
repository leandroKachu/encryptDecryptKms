package routers

import (
	"example/hello/src/decrypt"
	"net/http"
)

var getKeys = Route{
	URI:    "/get/key/{id}",
	Method: http.MethodGet,
	Func:   decrypt.DecryptKey,
}
