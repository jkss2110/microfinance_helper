package main

import (
	"fmt"
	"microfinance/models"
	"microfinance/shared"
	"net/http"
)

func main() {
	fmt.Println("Connection Open")
	models.DBConnection()
	shared.HandlerHttp(shared.ApiEndPoint)
	http.ListenAndServe(":5000", nil)
	defer models.Close()
	fmt.Println("Connection Closed")
}
