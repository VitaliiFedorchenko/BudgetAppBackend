package utils

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	w http.ResponseWriter
}

func NewResponse(w http.ResponseWriter) *Response {
	w.Header().Set("Content-Type", "application/json")
	return &Response{w: w}
}

func (r *Response) ResponseJSON(data interface{}, statusCode ...int) error {
	code := http.StatusOK
	if len(statusCode) > 0 {
		code = statusCode[0]
	}
	r.w.Header().Set("Content-Type", "application/json")
	r.w.WriteHeader(code)
	return json.NewEncoder(r.w).Encode(data)
}
