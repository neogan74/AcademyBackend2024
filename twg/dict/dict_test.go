package dict

import "testing"

func TestSearch(t *testing.T) {
	dictionary := map[string]string{"test": "this is just test"}

	got := Search(dictionary, "test")
	want := "this is just test"

	assertStrings(t, got, want)

}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
