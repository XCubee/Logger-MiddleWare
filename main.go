package main

import (
	"fmt"
	"net/http"
	"time"
)

func logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("started %s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) //Pass Middleware to next Middleware or final handler
		duration := time.Since(start)
		fmt.Printf("Completed %s %s in %v", r.Method, r.URL.Path, duration)
	})

}
func main() {

	http.Handle("/", logger(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(2 * time.Second)
		w.Write([]byte("Hello THERE"))
	})))

	http.ListenAndServe(":8080", nil)
}
