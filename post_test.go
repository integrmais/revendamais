package revendamais_test

import (
	"os"
	"testing"
	"net/http"
	"net/http/httptest"
	"github.com/integrmais/revendamais"
)

func TestListPosts(t *testing.T) {
	uriMock := "/application/index.php/apiGeneratorXml/generator/sitedaloja/hash.xml"

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uriExpected := r.URL.String()

		if uriExpected != uriMock {
			t.Errorf("expected %v, got %v", uriExpected, uriMock)
		}

		f, _ := os.ReadFile("./assets/posts_example.xml")

		w.Header().Add("Content-Type", "application/xml")
		w.Write([]byte(f))
	}))

	wp := revendamais.NewClient(
		server.URL,
		"hash",
	)

	posts, err := wp.Posts.List()
	if err != nil {
		t.Errorf("%v", err)
	}

	if len(posts.Posts) == 0 {
		t.Errorf("Got %v posts, expected 1", len(posts.Posts))
	}

	for _, post := range posts.Posts {
		if len(post.LargeImages) != 2 {
			t.Errorf("Expected 1 image, got %d", len(post.LargeImages))
		}
	}
}
