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

	"github.com/matt-wratt/go-chat/api"
)

var (
	binding    string
	port       int
	messageAPI = api.NewAPI()
)

func init() {
	flag.StringVar(&binding, "b", "0.0.0.0", "bind")
	flag.IntVar(&port, "p", 8080, "server port")
}

func main() {
	flag.Parse()

	os.Setenv("PORT", fmt.Sprintf("%d", port+1))

	if err := command("Building client", "client", "npm", "run", "build").Run(); err != nil {
		log.Fatal("Failed to build client")
	}

	if err := command("Starting client", "client", "npm", "run", "serve").Start(); err != nil {
		log.Fatalf("Failed to start client %v", err)
	}

	var proxy *httputil.ReverseProxy
	if clientURL, err := url.Parse(fmt.Sprintf("http://localhost:%d", port+1)); err == nil {
		proxy = httputil.NewSingleHostReverseProxy(clientURL)
	} else {
		log.Fatalf("Failed to start client %v", err)
	}

	host := fmt.Sprintf("%s:%d", binding, port)
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
		log.Fatalf("Server error: %v", err)
	}
	log.Printf("Stopping %s", host)
}

func command(description, path, name string, args ...string) *exec.Cmd {
	log.Printf(description)
	cmd := exec.Command(name, args...)
	cmd.Dir = filepath.Join(path)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd
}
