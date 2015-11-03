package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	"github.com/matt-wratt/messenger/api"
)

var messageAPI = api.NewAPI()

func main() {
	binding := flag.String("b", "0.0.0.0", "bind")
	port := flag.Int("p", 8080, "server port")
	flag.Parse()

	log.Printf("Starting client")
	path := filepath.Join("client")
	cmd := exec.Command("npm", "run", "serve")
	os.Setenv("PORT", fmt.Sprintf("%d", *port+1))
	cmd.Dir = path
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Printf("Failed to start client %v", err)
		return
	}

	var proxy *httputil.ReverseProxy
	if clientURL, err := url.Parse(fmt.Sprintf("http://localhost:%d", *port+1)); err == nil {
		proxy = httputil.NewSingleHostReverseProxy(clientURL)
	} else {
		log.Printf("Failed to create proxy %v", err)
		return
	}

	host := fmt.Sprintf("%s:%d", *binding, *port)
	log.Printf("Starting at %s", host)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if regexp.MustCompile("^/api").MatchString(path) {
			messageAPI.APIHandleFunc(w, r)
		} else {
			proxy.ServeHTTP(w, r)
		}
	})
	if err := http.ListenAndServe(host, nil); err != nil {
		log.Printf("%v", err)
	}
	log.Printf("Stopping %s", host)
}
