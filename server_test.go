package main

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("returns html with hello world", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/hello", nil)
		response := httptest.NewRecorder()

		helloHandler(response, request)

		got := response.Body.String()
		want := "Hello, world!"

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}

func TestForm(t *testing.T) {
	t.Run("returns html with form", func(t *testing.T) {
		data := url.Values{
			"name":    []string{"John"},
			"address": []string{"address"},
		}
		request, _ := http.NewRequest(http.MethodPost, "/form", strings.NewReader(data.Encode()))
		request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		response := httptest.NewRecorder()

		formHandler(response, request)

		got := response.Body.String()
		want := "Name: John, Address: address"
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
