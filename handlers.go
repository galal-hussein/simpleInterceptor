package main

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"html"
	"net/http"

	"github.com/Sirupsen/logrus"
)

// Print the API Request
func printAPIRequest(request Interceptor) {
	logrus.Debug("Request Intercepted.....")
	logrus.Debug("UUID: ", request.UUID)
	logrus.Debug("Headers: ", request.Headers)
	logrus.Debug("Body: ", request.Body)
	logrus.Debug("API Path: ", request.APIPath)
	logrus.Debug("Request Ended.....")
}

// Index route
func Index(w http.ResponseWriter, r *http.Request) {
	var request Interceptor
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		logrus.Fatal("Error: ", err)
		return
	}
	logrus.Infof("Endpoint Invoked %q", html.EscapeString(r.URL.Path))
	printAPIRequest(request)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// Secret route
func Secret(w http.ResponseWriter, r *http.Request) {
	logrus.Infof("Endpoint Invoked %q", html.EscapeString(r.URL.Path))

	var request Interceptor
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		logrus.Fatal("Error: ", err)
		return
	}

	// mac write the content
	key := []byte("rancher123")
	bodyContent, err := json.Marshal(r.Body)
	if err != nil {
		logrus.Fatal("Error: ", err)
		return
	}

	expectedSignature := signMessage(bodyContent, key)
	logrus.Debugf("Signature generated: %v", expectedSignature)
	existingSignature := r.Header.Get("X-API-Auth-Signature")
	logrus.Debugf("Existing Signature: %v", existingSignature)

	if hmac.Equal([]byte(existingSignature), []byte(expectedSignature)) {
		logrus.Infof("Signature Verified...")
	} else {
		logrus.Fatal("Error: Signature not verified")
		return
	}
	printAPIRequest(request)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	logrus.Infof("Endpoint Ended")
}

func signMessage(body []byte, key []byte) string {
	// A known secret key
	mac := hmac.New(sha512.New, key)
	mac.Write(body)
	signature := mac.Sum(nil)
	encodedSignature := base64.URLEncoding.EncodeToString(signature)
	return encodedSignature
}
