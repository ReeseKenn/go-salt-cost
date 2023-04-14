package controller_test

import (
	"net/http"
	"net/http/httptest"
	"salt-cost/internal/router"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPing(t *testing.T) {
	r := router.New()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong\n", w.Body.String())
}
