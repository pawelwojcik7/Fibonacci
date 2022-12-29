package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/fib", f)
	http.ListenAndServe(":8080", nil)
}
func f(w http.ResponseWriter, r *http.Request) {
	s := r.URL.Query().Get("n")
	n, e := strconv.Atoi(s)
	if e != nil {
		http.Error(w, "Invalid parameter", http.StatusBadRequest)
		return
	}
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%d", a)
	fmt.Fprintf(w, "%s", "\n Pawel Wojcik")
	fmt.Fprintf(w, "%s", "\n Fibonacci")
	fmt.Fprintf(w, "%s", "\n 2.1.5")
}
