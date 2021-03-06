package request

import (
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	assert.Equal(t, "Not Found", Error(404).Error())
	assert.Equal(t, "Bad Request", Error(400).Error())
}

func TestIsNotFound(t *testing.T) {
	assert.True(t, IsNotFound(Error(404)))
	assert.False(t, IsNotFound(Error(500)))
}

func TestIsStatus(t *testing.T) {
	assert.True(t, IsStatus(Error(404), 404))
	assert.True(t, IsStatus(Error(500), 500))
	assert.False(t, IsStatus(Error(500), 400))
}

func TestIsClient(t *testing.T) {
	assert.True(t, IsClient(Error(400)))
	assert.False(t, IsClient(Error(500)))
}

func TestIsServer(t *testing.T) {
	assert.True(t, IsServer(Error(500)))
	assert.False(t, IsServer(Error(400)))
}

func TestParam(t *testing.T) {
	u := &url.URL{RawQuery: "hello=world&bye=planet"}
	r := &http.Request{URL: u}
	assert.Equal(t, Param(r, "hello"), "world")
	assert.Equal(t, Param(r, "bye"), "planet")
}

func TestParamDefault(t *testing.T) {
	u := &url.URL{}
	r := &http.Request{URL: u}
	assert.Equal(t, ParamDefault(r, "hello", "oops"), "oops")
}
