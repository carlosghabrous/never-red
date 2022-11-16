package main

import "net/http"

//TODO: what if there are multiple allow methods?
func validMethod(apiMethod, requestMethod string, w http.ResponseWriter) bool {
	if apiMethod != requestMethod {
		w.Header().Set("Allow", apiMethod)
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
		return false
	}

	return true
}
