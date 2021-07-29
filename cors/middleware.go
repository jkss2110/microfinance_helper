package cors

import (
	"fmt"
	"net/http"
	"time"
)

func MiddlewareHandler(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Before handler")
		start := time.Now()
		handler.ServeHTTP(w, r)
		fmt.Printf("middleware finished %s", time.Since(start))
	})
}
