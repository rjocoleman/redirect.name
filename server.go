package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"golang.org/x/net/publicsuffix"
)

var defaultRecord = os.Getenv("DEFAULT_RECORD")

func fallback(w http.ResponseWriter, r *http.Request, reason string) {
	location := "http://redirect.name/"
	if reason != "" {
		location = fmt.Sprintf("%s#reason=%s", location, url.QueryEscape(reason))
	}
	http.Redirect(w, r, location, 302)
}

func hostnameLookup(host string) ([]string, error) {
	hostname := fmt.Sprintf("_redirect.%s", host)
	return net.LookupTXT(hostname)
}

func handler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(r.Host, ":")
	host := parts[0]

	txt, err := hostnameLookup(host)
	if err != nil {
		if defaultRecord != "" {
			tld, _ := publicsuffix.EffectiveTLDPlusOne(host)
			recursiveHost := fmt.Sprintf("%s.%s", defaultRecord, tld)

			txt, err = hostnameLookup(recursiveHost)
		}
	}

	if err != nil {
		fallback(w, r, fmt.Sprintf("Could not resolve hostname (%v)", err))
		return
	}

	for _, record := range txt {
		redirect := Translate(r.URL.String(), Parse(record))
		if redirect != nil {
			http.Redirect(w, r, redirect.Location, redirect.Status)
			return
		}
	}

	fallback(w, r, "No paths matched")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
	}

	http.HandleFunc("/", handler)
	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}

	log.Printf("Listening on http://127.0.0.1:%s", port)
	log.Fatal(srv.ListenAndServe())
}
