package routes

import (
	"fmt"

	"github.com/a-h/templ"
)

const (
	Home        = "/"
	Trips       = "/trips"
	TripBuilder = "/tripbuilder"
	NewTrips    = "/trips/new"
	List        = "/trips/list"
)

type QueryParams map[string]string

func SafeQueryParamBuilder(url string, params QueryParams) templ.SafeURL {
	url = fmt.Sprintf("%s?", url)
	for k, v := range params {
		url = fmt.Sprintf("%s&%s=%s", url, k, v)
	}
	return templ.URL(url)
}

func QueryParamBuilder(url string, params QueryParams) string {
	url = fmt.Sprintf("%s?", url)
	for k, v := range params {
		url = fmt.Sprintf("%s&%s=%s", url, k, v)
	}
	return url
}
