package main

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		defination := "this is just a test"
		err := dictionary.Add(word, defination)
		assertError(t, err, nil)
		assertDefinition(t, dictionary, "test", defination)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")
		assertError(t, err, ErrWordExist)
		assertDefinition(t, dictionary, word, defination)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is just a test"
		dictionary := Dictionary{word: defination}
		newDefination := "new defination"
		err := dictionary.Update(word, newDefination)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefination)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defination := "this is just a test"

		err := dictionary.Update(word, defination)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test defination"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})

	t.Run("non-exist word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}

func assertStrings(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
