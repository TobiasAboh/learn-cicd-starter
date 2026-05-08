package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "")
	_, err := GetAPIKey(h)
	if !reflect.DeepEqual(nil, err) {
		t.Fatalf("expected: %v, got: %v", nil, err)
	}
}

func TestGetAPIKeyValue(t *testing.T) {
	h := http.Header{}
	h.Set("Authorization", "ApiKey Tobias")
	got, _ := GetAPIKey(h)
	want, _ := "Aboh", ""
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
