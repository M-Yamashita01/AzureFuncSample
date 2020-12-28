package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/M-Yamashita01/AzureFuncSample/hi"
)

func main() {
	port, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
	if !exists {
		port = "5000"
	}

	envTest := os.Getenv("ENV_TEST")
	fmt.Println("Env: " + envTest)

	http.HandleFunc("/api/hi", hi.ExportParticipants)
	fmt.Println("Go server listening on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))

}
