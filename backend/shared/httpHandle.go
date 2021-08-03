package shared

import (
	"microfinance/controller"
	"microfinance/cors"
	"net/http"
)

func HandlerHttp(apiEndpoint string) {
	loanListHandler := http.HandlerFunc(controller.HandleHttp)
	loanItemHandler := http.HandlerFunc(controller.HandleSingleHttp)

	// Adding Middleware for before and after calculations
	http.Handle(apiEndpoint+"/loandetails", cors.MiddlewareHandler(loanListHandler))
	http.Handle(apiEndpoint+"/loandetails/", cors.MiddlewareHandler(loanItemHandler))
}
