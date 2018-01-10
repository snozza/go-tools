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
	assert.Equal(t, "<User>\n  <First>Tobi</First>\n  <Last>Ferret</Last>\n</User>", string(res.Body.Bytes()))
	assert.Equal(t, "application/xml", res.HeaderMap["Content-Type"][0])
}
