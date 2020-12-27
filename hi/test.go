package hi

import (
	"fmt"
	"net/http"
	"os"
)

func ExportParticipants(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "EnvironmentTest="+os.Getenv("SLACK_SIGNING_SECRET"))
	fmt.Fprintf(w, "EnvironmentTes="+os.Getenv("ENV_TEST"))
}
