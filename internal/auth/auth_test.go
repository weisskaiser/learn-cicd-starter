package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {

	want := "John"
	given := http.Header{}
	given.Set("Authorization", "ApiKey John")

	got, err := GetAPIKey(given)

	if err != nil {
		t.Fatalf("shouldn't have errors %v", err)
	}

	if got != want {
		t.Fatalf("want %v got %v", want, got)
	}
}

func TestGetInvalidAPIKey(t *testing.T) {
	given := http.Header{}
	given.Set("Authorization", "Bearer John")

	_, err := GetAPIKey(given)

	if err == nil {
		t.Fatal("should have errors")
	}

	t.Log(err.Error())
}
