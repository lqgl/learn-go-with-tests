package maps

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		// dictionary := map[string]string{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertString(t, got, want)
	})

	t.Run("unknow word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %s, but want %s", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error '%s' want '%s'", got, want)
	}

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{}
	dictionary.Add("test", "this is just a test")
	want := "this is just a test"
	got, err := dictionary.Search("test")

	if err != nil {
		t.Fatal("should find added word:", err)
	}

	if want != got {
		t.Errorf("got '%s', but want '%s'", got, want)
	}
}
