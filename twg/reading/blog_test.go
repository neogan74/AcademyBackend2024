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
	got := posts[0]
	want := Post{Title: "Post 1"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v posts, wanted %+v posts", got, want)
	}
}
