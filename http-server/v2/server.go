package v1

import (
	"fmt"
	"net/http"
)

func PlayerServer(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "20")
}
