package client

import (
	"net/http"
	"net/http/cookiejar"
)
func NewClient() *http.Client{
	jar,_:=cookiejar.New(nil)
	return &http.Client{Jar:jar}
}