package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const versionENV = "VERSION"

type service struct {
	version string
}

func (s *service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set(versionENV, s.version)

	for s, strs := range r.Header {
		w.Header()[s] = strs
	}

	fmt.Printf("remote IP: %s", strings.Split(r.RemoteAddr, ":")[0])
}

func main() {
	ver := os.Getenv(versionENV)

	s := &service{version: ver}

	http.Handle("/healthz", s)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
