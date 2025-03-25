package concurency

import (
	"reflect"
	"testing"
)

func mockWebsiteChecker(url string) bool {
	if url == "http://badwebsite.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"http://goodwebsite.com",
		"http://anothergoodwebsite.com",
		"http://badwebsite.com",
	}

	want := map[string]bool{
		"http://goodwebsite.com":        true,
		"http://anothergoodwebsite.com": true,
		"http://badwebsite.com":         false,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %v, want %v", got, want)
	}
}
