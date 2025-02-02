package reading

import (
	"reflect"
	"testing"
	"testing/fstest"
)

func TestNewBlogPosts(t *testing.T) {
	const (
		firstBody = `Title: Post 1
Description: Description 1
Tags: tss, go
---
Hello
World`
		secondBody = `Title: Post 2
Description: Description 2
Tags: rust, borrow-checker
---
Certificed
K8s
Specialist`
	)
	fs := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello-world2.md": {Data: []byte(secondBody)},
	}
	posts, err := NewPostsFromFS(fs)
	if err != nil {
		t.Error("Cannot read posts")
	}

	assertPost(t, posts[0], Post{Title: "Post 1",
		Description: "Description 1",
		Tags:        []string{"tss", "go"},
		Body: `Hello
World`,
	})
}

func assertPost(t *testing.T, got Post, want Post) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v posts, wanted %+v posts", got, want)
	}
}
