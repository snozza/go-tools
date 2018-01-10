package response

import (
	"encoding/xml"
	"net/http"
)

// XML response with optional status Code
func XML(w http.ResponseWriter, val interface{}, code ...int) {
	var b []byte
	var err error

	if Pretty {
		b, err = xml.MarshalIndent(val, "", " ")
	} else {
		b, err = xml.Marshal(val)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if len(code) > 0 {
		w.WriteHeader(code[0])
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(b)
}
