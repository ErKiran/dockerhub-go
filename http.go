package dockerhub

import (
	"fmt"
	"net/http"
)

// CheckResponse checks a given HTTP response for errors and returns
// them if present.
func CheckResponse(r *http.Response) error {
	status := r.StatusCode
	if status >= 200 && status <= 299 {
		return nil
	}

	return fmt.Errorf("Request failed with status %d", status)
}
