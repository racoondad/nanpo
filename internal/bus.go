package internal

import "net/http"

// Bus Message bus, using http header to pass through data.
type Bus struct {
	http.Header
}
type BusHandler func(Worker)
