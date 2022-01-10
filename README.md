# reverse-proxy-go
A Simple And Stupid Reverse Proxy With Golang


# Usage
````bash
./reverse-proxy --help

            This is a simple Http reverse proxy service.

            When you specify the client address, use the prefix /fe to access it.
            When you specify the server side address, use the prefix /be to access it.

Usage:
  reverse-proxy [flags]

Flags:
      --clientURL string   Client URL
  -h, --help               help for reverse-proxy
      --port int           Proxy Port (default 8080)
      --serverURL string   Server URL
````
