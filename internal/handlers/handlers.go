// SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
//
// SPDX-License-Identifier: Apache-2.0

package handlers

import "net/http"

// NotFound is a generic handler that writes a NotFound response to the response writer.
func NotFound() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"code":404,"message":"not found"}`)) //nolint:errcheck,gosec
	})
}

// Ping is a handler that writes 200 OK to the response writer.
func Ping() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"code":200,"message":"ok"}`)) //nolint:errcheck,gosec
	})
}
