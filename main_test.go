package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	cases := []struct {
		in, out string
	}{
		{"http://localhost:8080/test@golang.org", "gopher"},
		{"http://localhost:8080/test@golang.or", "Hello, world!!"},
	}
	for _, c := range cases {
		req, err := http.NewRequest(http.MethodGet, c.in, nil)
		if err != nil {
			t.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		handlerRegex(rec, req)

		if rec.Code != http.StatusOK {
			t.Errorf("expected status 200; got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), c.out) {
			t.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	}

}

func BenchmarkHandler(b *testing.B) {
	for i := 0; i < b.N; i++ {
		req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/test@golang.org", nil)
		if err != nil {
			b.Fatalf("could not create request: %v", err)
		}
		rec := httptest.NewRecorder()
		handlerRegex(rec, req)

		if rec.Code != http.StatusOK {
			b.Errorf("expected status 200; got %d", rec.Code)
		}
		if !strings.Contains(rec.Body.String(), "gopher") {
			b.Errorf("unexpected body in response: %q", rec.Body.String())
		}
	}
}
