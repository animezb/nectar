package hummingbird

import (
	"net/http"
)

const userAgent = "Mozilla/5.0 (Nectar Hummingbird Client; http://github.com/animezb/nectar"

var defaultClient *http.Client
var cookieJar http.CookieJar

type apiError struct {
	Msg string `json:"error"`
}

func (e apiError) Error() string {
	return e.Msg
}

func (e apiError) String() string {
	return e.Msg + "ff"
}

func init() {
	defaultClient = &http.Client{Jar: cookieJar}
}
