package dockerhub

import (
	"fmt"
	"net/http"
)

// checkResponse checks a given HTTP response for errors and returns
// them if present.
func checkResponse(r *http.Response) error {
	status := r.StatusCode
	if status >= 200 && status <= 299 {
		return nil
	}

	return fmt.Errorf("request failed with status %d", status)
}
