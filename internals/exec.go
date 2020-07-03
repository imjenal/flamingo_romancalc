package internals

import "net/http"

type Tasks interface {
	Execute(http.ResponseWriter, string, string)
}

