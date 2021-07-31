package main

import (
	"microfinance/shared"
	"net/http"
)

func main() {
	shared.DBConnection()
	shared.HandlerHttp(shared.ApiEndPoint)
	http.ListenAndServe(":5000", nil)
}
