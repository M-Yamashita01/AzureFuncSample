// package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
// 	"os"

// 	"github.com/M-Yamashita01/AzureFuncSample/hi"
// )

// func main() {
// 	// port, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
// 	// if !exists {
// 	port := "8888"
// 	// }

// 	envTest := os.Getenv("ENV_TEST")
// 	fmt.Println("Env: " + envTest)

// 	http.HandleFunc("/api/hi", hi.ExportParticipants)
// 	fmt.Println("Go server listening on port", port)
// 	log.Fatal(http.ListenAndServe(":"+port, nil))

// }

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"
	name := r.URL.Query().Get("name")
	if name != "" {
		message = fmt.Sprintf("Hello, %s. This HTTP triggered function executed successfully.\n", name)
	}
	fmt.Fprint(w, message)
	hoge := os.Getenv("ENV_TEST")
	log.Printf("ENV hoge is " + hoge)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/HttpTrigger1", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
