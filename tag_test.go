package dockerhub

import (
	"context"
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestTagService_GetTags(t *testing.T) {
	client, mux, teardown := makeMockClient()
	defer teardown()

	tags := Tags{}

	namespace := "namespace"
	repo := "repo"
	page := 100

	uri := fmt.Sprintf("/repositories/%s/%s/tags/?page_size=%d", namespace, repo, page)

	mux.HandleFunc(uri, func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, r, http.MethodGet)
		w.WriteHeader(http.StatusOK)
		w.Write(mustJSONMarshal(tags))
	})

	res, err := client.Tag.GetTags(context.Background(), namespace, repo, page)
	if err != nil {
		t.Errorf("Tag.GetTags returned error: %v", err)
	}

	if !reflect.DeepEqual(res, tags) {
		t.Errorf("Tags is %v; want %v", res, tags)
	}
}
