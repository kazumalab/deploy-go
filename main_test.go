package main

import (
  "testing"
  "net/http/httptest"
)

func TestHandle(t *testing.T) {
  req := httptest.NewRequest("GET", "/", nil)
  w := httptest.NewRecorder()
  handle(w, req)

  resp := w.Result()

  if resp.StatusCode != 200 {
    t.Fatal("Error")
  }
}
