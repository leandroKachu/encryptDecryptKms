package routers

import (
	"example/hello/src/encrypt"
	"net/http"
)

var routerCreateKey = []Route{
	{
		URI:    "/create/key",
		Method: http.MethodPost,
		Func:   encrypt.EncryptKey,
	},
}
