package server

import (
    "fmt"
    "github.com/gorilla/mux"
    "log"
    "net/http"
    "net/http/httputil"
    "net/url"
)

var prefixLength int

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
    r.URL.Path = fmt.Sprintf("%s", r.URL.Path[prefixLength:])
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

func Start(urls map[string]string, port int) {
    r := mux.NewRouter()
    for prefix, url := range urls {
	proxy, err := NewProxy(fmt.Sprintf("%s", url))
	if err != nil {
	    panic(err)
	}
	prefixLength = len(prefix)
	r.PathPrefix(fmt.Sprintf("%s", prefix)).HandlerFunc(ProxyRequestHandler(proxy))
    }

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
