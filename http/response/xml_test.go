package response

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXMLPretty(t *testing.T) {
	Pretty = true
	res := httptest.NewRecorder()
	XML(res, &User{"Tobi", "Ferret"})
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, "<User>\n <First>Tobi</First>\n <Last>Ferret</Last>\n</User>", string(res.Body.Bytes()))
	assert.Equal(t, "application/xml", res.HeaderMap["Content-Type"][0])
}

func TestXML(t *testing.T) {
	Pretty = false
	res := httptest.NewRecorder()
	XML(res, &User{"Tobi", "Ferret"})
	assert.Equal(t, 200, res.Code)
	assert.Equal(t, `<User><First>Tobi</First><Last>Ferret</Last></User>`, string(res.Body.Bytes()))
	assert.Equal(t, "application/xml", res.HeaderMap["Content-Type"][0])
}

func TestXMLError(t *testing.T) {
	res := httptest.NewRecorder()
	invalidBody := make(chan int, 0)
	XML(res, invalidBody)
	assert.Equal(t, 500, res.Code)
	assert.Equal(t, "text/plain; charset=utf-8", res.HeaderMap["Content-Type"][0])
}

func TestXMLCode(t *testing.T) {
	res := httptest.NewRecorder()
	XML(res, &User{}, 202)
	assert.Equal(t, 202, res.Code)
}
