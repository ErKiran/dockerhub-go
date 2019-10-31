package dockerhub

import (
	"context"
	"net/http"
	"testing"
)

func TestAuthService_Login(t *testing.T) {
	client, mux, teardown := makeMockClient()
	defer teardown()

	username := "username"
	password := "password"
	token := "bogus"

	mux.HandleFunc("/users/login/", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, r, http.MethodPost)
		assertNoHeader(t, r, "Authorization")
		assertBody(t, r, string(mustJSONMarshal(&LoginRequest{
			Username: username,
			Password: password,
		})))
		w.WriteHeader(http.StatusCreated)
		w.Write(mustJSONMarshal(&LoginResponse{
			Token: token,
		}))
	})

	if err := client.Auth.Login(context.Background(), username, password); err != nil {
		t.Errorf("Auth.Login returned error: %v", err)
	}

	if got := client.authToken; got != token {
		t.Errorf("client.authToken is %s; want %s", got, token)
	}
}

func TestAuthService_Login_NoToken(t *testing.T) {
	client, mux, teardown := makeMockClient()
	defer teardown()

	mux.HandleFunc("/users/login/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(""))
	})

	err := client.Auth.Login(context.Background(), "username", "password")
	if err == nil {
		t.Errorf("Auth.Login succeeded without getting token response")
	}

	want := "Did not recieve token"
	if got := err.Error(); got != want {
		t.Errorf(`Auth.Login error "%s"; want "%s"`, got, want)
	}
}
