package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/msg", msgHandler)
	http.HandleFunc("/", mainHandler)
	http.ListenAndServe(":80", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	out, _ := ioutil.ReadFile("msg.json")
	//json response
	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

func msgHandler(w http.ResponseWriter, r *http.Request) {
	//check type of request
	if r.Method == "POST" {
		// Get value from request
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
			return
		}
		// Write response
		w.Write(body)
		ioutil.WriteFile("msg.json", body, 0644)
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
