package main

import (
	"fmt"
)

import "os"
import "encoding/json"
import "net/http"

func main() {
	http.HandleFunc("/", HttpHeaders)
	_ = http.ListenAndServe(":8080", nil)
}

func HttpHeaders(w http.ResponseWriter, r *http.Request) {
	body := make(map[string]interface{})

	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)

	body["a_path"] = r.RequestURI
	body["a_hostname"], _ = os.Hostname()

	body["b_headers"] = r.Header

	body["c_uid"] = os.Getuid()
	body["c_gid"] = os.Getgid()
	body["d_euid"] = os.Geteuid()
	body["d_egid"] = os.Getegid()

	body["z_environment"] = os.Environ()

	if reqHeadersBytes, err := json.Marshal(body); err != nil {
	} else {
		_, _ = fmt.Fprintf(w, string(reqHeadersBytes))
	}

}
