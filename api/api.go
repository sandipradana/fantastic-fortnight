package api

import (
	"fantastic-fortnight/backend/cmd/vercel"
	"net/http"
)

func Api(httpResponse http.ResponseWriter, httpRequest *http.Request) {
	vercel.Handler(httpResponse, httpRequest)
}
