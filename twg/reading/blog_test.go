package reading

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello-world2.md": {Data: []byte("Title: Post 2")},
	}
	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Error("Cannot read posts")
	}

	assertPost(t, posts[0], Post{Title: "Post 1"})
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v posts, wanted %+v posts", got, want)
	}
}
