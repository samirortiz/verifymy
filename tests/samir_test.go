package tests

import (
	"net/http"
	"testing"

	"github.com/steinfletcher/apitest"
)

func TestGetUser(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		msg := `{"name": "samir"}`
		_, _ = w.Write([]byte(msg))
		w.WriteHeader(http.StatusOK)
	}

	apitest.New(). // configuration
			HandlerFunc(handler).
			Get("/users"). // request
			Expect(t).     // expectations
			Body(`{"name": "samir"}`).
			Status(http.StatusOK).
			End()
}

func TestGetUser_NotFound(t *testing.T) {
	apitest.New().
		Handler(newRouter()).
		Get("/users/1515").
		Expect(t).
		Status(http.StatusNotFound).
		End()
}
