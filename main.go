package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

const (
	DEFAULT_PORT              = "8080"
	CF_FORWARDED_URL_HEADER   = "X-Cf-Forwarded-Url"
	CF_PROXY_SIGNATURE_HEADER = "X-Cf-Proxy-Signature"
)

func main() {
	var port string
	if port = os.Getenv("PORT"); len(port) == 0 {
		port = DEFAULT_PORT
	}

	log.SetOutput(os.Stdout)

	proxy := NewProxy()

	log.Fatal(http.ListenAndServe(":"+port, proxy))
}

func NewProxy() http.Handler {
	reverseProxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			forwardedURL := req.Header.Get(CF_FORWARDED_URL_HEADER)

			log.Println("Received request: ")
			log.Printf("%s: %s\n", CF_FORWARDED_URL_HEADER, forwardedURL)
			log.Printf("%s: %s\n", CF_PROXY_SIGNATURE_HEADER, req.Header.Get(CF_PROXY_SIGNATURE_HEADER))
			log.Println("")
			log.Printf("Headers: %#v\n", req.Header)
			log.Println("")

			body, err := ioutil.ReadAll(req.Body)
			if err != nil {
				log.Fatalln(err.Error())
			}
			log.Printf("Request Body: %s\n", string(body))

			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			// Note that url.Parse is decoding any url-encoded characters.
			url, err := url.Parse(forwardedURL)
			if err != nil {
				log.Fatalln(err.Error())
			}

			log.Printf("Forwarding to: %s\n", forwardedURL)

			req.URL = url
			req.Host = url.Host
		},
	}
	return reverseProxy
}
