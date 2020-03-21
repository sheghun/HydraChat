package main

import (
	"Hydra/hlogger"
	"fmt"
	"net/http"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Print("Starting Hydra web service")
	http.HandleFunc("/", sroot)
	http.ListenAndServe(":8080", nil)
}

func sroot(w http.ResponseWriter, r *http.Request) {
	logger := hlogger.GetInstance()
	fmt.Fprint(w , "Welcome to the hydra software system")

	logger.Println("Received and http request on the root url")
}
