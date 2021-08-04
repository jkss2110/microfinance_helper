package shared

import (
	"microfinance/controller"
	"microfinance/cors"
	"microfinance/models"
	"net/http"

	"golang.org/x/net/websocket"
)

func HandlerHttp(apiEndpoint string) {
	loanListHandler := http.HandlerFunc(controller.HandleHttp)
	loanItemHandler := http.HandlerFunc(controller.HandleSingleHttp)

	// Adding Middleware for before and after calculations
	http.Handle(apiEndpoint+"/loandetails", cors.MiddlewareHandler(loanListHandler))
	http.Handle(apiEndpoint+"/loandetails/", cors.MiddlewareHandler(loanItemHandler))

	// Websocket
	http.Handle(apiEndpoint+"/websocket", websocket.Handler(models.LoanDashboardSocket))
}
