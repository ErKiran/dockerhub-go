package dockerhub

// String returns a pointer to a string for configuration.
func String(s string) *string {
	return &s
}

// StringValue returns the value of a String pointer
func StringValue(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
