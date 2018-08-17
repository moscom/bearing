package bearing

import "net/http"

// Context object
type Context struct {
	Request *http.Request
}
