package webclient

import (
	"context"
	"net/http"
	"net/url"
)

type Webclient interface {
	Header(header http.Header) Webclient
	Url(url string) Webclient
	Query(values url.Values) Webclient

	Model(model interface{}) Webclient
	StatusCode(statCode *int) Webclient

	Get(ctx context.Context) Webclient
	Post(ctx context.Context) Webclient

	Response() http.Response
	Error() error
	Return() (http.Response, error)
}
