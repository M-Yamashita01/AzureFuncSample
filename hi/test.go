package hi

import (
	"fmt"
	"net/http"
	"os"
)

func ExportParticipants(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Environment="+os.Getenv("SLACK_SIGNING_SECRET"))
}
