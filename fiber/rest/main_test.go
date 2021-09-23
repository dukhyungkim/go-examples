package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http/httptest"
	"testing"
)

func Test_Handler(t *testing.T) {
	app := setup()

	resp, err := app.Test(httptest.NewRequest("GET", "/hello", nil))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	assert.Equal(t, 200, resp.StatusCode)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	defer resp.Body.Close()
	assert.Equal(t, "Hello, World 👋!", string(body))
}
