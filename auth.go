package dockerhub

import (
	"context"
	"errors"
	"net/http"
)

// AuthService handles communication with the auth related
// methods of the Dockerhub API.
type AuthService service

// LoginRequest represents the payload to be sent to login to the
// Dockerhub API.
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoginResponse represents the payload to be responded to a successful
// Dockerhub API login request.
type LoginResponse struct {
	Token string `json:"token"`
}

// Login authenticates with the Dockerhub API with the given given
// username and password.
func (s *AuthService) Login(ctx context.Context, username, password string) error {
	p := &LoginRequest{username, password}
	req, err := s.client.NewRequest(http.MethodPost, "/users/login/", p)
	if err != nil {
		return err
	}

	res := &LoginResponse{}
	if _, err := s.client.Do(ctx, req, res); err != nil {
		return err
	}

	if len(res.Token) == 0 {
		return errors.New("did not recieve token")
	}

	s.client.SetAuthToken(res.Token)
	return nil
}
