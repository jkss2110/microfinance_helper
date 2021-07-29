package main

import (
	"microfinance/shared"
	"net/http"
)

func main() {
	shared.HandlerHttp()
	http.ListenAndServe(":5000", nil)
}
