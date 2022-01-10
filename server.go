package main

import (
    "fmt"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

func NewProxy(targetHost string) (*httputil.ReverseProxy, error) {
    url, err := url.Parse(targetHost)
    if err != nil {
	return nil, err
    }

    proxy := httputil.NewSingleHostReverseProxy(url)

    originalDirector := proxy.Director
    proxy.Director = func(req *http.Request) {
	originalDirector(req)
	modifyRequest(req)
    }

    proxy.ModifyResponse = modifyResponse()
    proxy.ErrorHandler = errorHandler()
    return proxy, nil
}

func modifyRequest(r *http.Request) {
    r.Header.Set("X-Proxy", "Simple-Reverse-Proxy")
}

func errorHandler() func(http.ResponseWriter, *http.Request, error) {
    return func(w http.ResponseWriter, r *http.Request, err error) {
	fmt.Printf("Got error while modifying response: %v \n", err)
	return
    }
}

func modifyResponse() func(*http.Response) error {
    return func(res *http.Response) error {
	res.Header.Set("X-Server-Proxy", "Simple-Reverse-Proxy")
	return nil
    }
}

func ProxyRequestHandler(proxy *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
    return func(w http.ResponseWriter, r *http.Request) {
	proxy.ServeHTTP(w, r)
    }
}

func main() {
    proxy, err := NewProxy("http://localhost:8888/")
    if err != nil {
	panic(err)
    }

    http.HandleFunc("/", ProxyRequestHandler(proxy))
    log.Fatal(http.ListenAndServe(":8080", nil))
}
