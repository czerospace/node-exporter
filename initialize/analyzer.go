package initialize

import (
	"fmt"
	"net/http"
)

func handleHealthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "OK")
}

func init() {
	http.HandleFunc("/health", handleHealthcheck)
}
