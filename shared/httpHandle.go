package shared

import (
	"fmt"
	"microfinance/controller"
	"net/http"
	"time"
)

func middlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finished %s", time.Since(start))
	})
}
func HandlerHttp() {
	loanListHandler := http.HandlerFunc(controller.HandleHttp)
	loanItemHandler := http.HandlerFunc(controller.HandleSingleHttp)

	// Adding Middleware for before and after calculations
	http.Handle("/loandetails", middlewareHandler(loanListHandler))
	http.Handle("/loandetails/", middlewareHandler(loanItemHandler))
}
