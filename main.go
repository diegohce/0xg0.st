package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/golang/glog"
)

// Dead simple router that just does the **perform** the job
func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case strings.Contains(r.Header.Get("Content-type"), "multipart/form-data"):
		upload(w, r)
		break
	case uuidMatch.MatchString(r.URL.Path):
		getFile(w, r)
		break
	default:
		http.ServeFile(w, r, "./templates/index.html")
		break
	}
}

// Route handling, logging and application serving
func main() {
	// Random seed creation
	rand.Seed(time.Now().Unix())

	// Flags for the leveled logging
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: 0xg0.st -stderrthreshold=[INFO|WARNING|FATAL] -log_dir=[string]\n")
		flag.PrintDefaults()
		os.Exit(2)
	}

	flag.Parse()
	glog.Flush()

	// Routing
	http.HandleFunc("/", router)
	http.ListenAndServe(":8000", nil)
}
