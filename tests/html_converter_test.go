package tests

import (
	"github.com/gavv/httpexpect/v2"
	"html-convert/internal/http-server/handlers/convert"
	"net/url"
	"testing"
)

const (
	host = "localhost:8081"
)

func TestHtmlConvert(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}
	var e = httpexpect.Default(t, u.String())
	e.POST("/convert").
		WithJSON(convert.Request{
			Html: "<html><body>Hello</body></html>",
		}).
		Expect().
		Status(200).
		HasContentType("application/pdf").
		Body().
		Raw()
}
