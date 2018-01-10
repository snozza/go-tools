package response

import (
	"encoding/xml"
	"net/http"
)

// XML response with optional status Code
func XML(w http.ResponseWriter, val interface{}, code ...int) {
	var b []byte
	//var err error

	if Pretty {
		b, _ = xml.MarshalIndent(val, "", " ")
	}

	w.Header().Set("Content-Type", "application/xml")

	w.Write(b)
}
