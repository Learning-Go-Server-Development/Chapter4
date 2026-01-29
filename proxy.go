package goproxy

import "net/http"

// Proxy Proxy
type Proxy interface {
	Do(req *http.Request, obj any) (bool, int)
}

// go mod init github.com/Learning-Go-Server-Development/Chapter4
