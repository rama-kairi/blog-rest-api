package utils

import (
	"net/http"
	"strconv"
	"strings"
)

// Enum for HTTP methods
type HttpMethod string

const (
	GET    HttpMethod = "GET"
	POST   HttpMethod = "POST"
	PUT    HttpMethod = "PUT"
	DELETE HttpMethod = "DELETE"
)

// CheckMethod - Check if the HTTP method is the same as the one passed in
func CheckMethod(httpMethod string, checkMethod HttpMethod) bool {
	return httpMethod == string(checkMethod)
}

// GetUrlParamId - Get the id from the url query parameters
func GetUrlParamId(r *http.Request) (int, error) {
	path := r.URL.Path
	isSTr := strings.Split(path, "/")
	id, err := strconv.Atoi(isSTr[len(isSTr)-1])
	return id, err
}

// GetQueryParams - Get the query parameters from the url
func GetQueryParams(r *http.Request) map[string]string {
	queryParams := make(map[string]string)
	for key, value := range r.URL.Query() {
		queryParams[key] = strings.Join(value, "")
	}
	return queryParams
}

// SuccessResponse - This function writes a success response to the response writer
func SuccessResponse(w http.ResponseWriter, status int, data []byte) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}
