package handler_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/infysumanta/go-hello-world-server/internal/handler"
)

func TestHelloWorldHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(handler.HelloWorldHandler)

    handler.ServeHTTP(rr, req)

    if rr.Code != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", rr.Code, http.StatusOK)
    }

    expected := "Hello World\n"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}