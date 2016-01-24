package main

import (
	"bufio"
	"crypto/sha1"
	"encoding/base64"
	"net/http"
	"os"
	"strings"
)

func basicAuthentication(reqHandler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if config.AuthEnabled {
			username, password, ok := r.BasicAuth()
			if !ok {
				renderUnauthorized(w, "Authentication failed - Provide Basic Authentication - username:password")
				return
			}
			if !validate(username, password) {
				renderUnauthorized(w, "Authentication failed - incorrect username or password")
				return
			}
		}
		reqHandler(w, r)
	}
}

func validate(username, password string) bool {
	file, err := os.Open(config.PasswdFilePath)
	if err != nil {
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		params := strings.Split(scanner.Text(), ":")
		if username == params[0] && matchPassword(password, params[1]) {
			return true
		}
	}
	return false
}

func matchPassword(savedPwd string, sentPwd string) bool {
	hash := sha1.New()
	hash.Write([]byte(savedPwd))
	pwdCheck := strings.Replace(base64.URLEncoding.EncodeToString(hash.Sum(nil)), "-", "+", -1)
	pwdCheck = strings.Replace(pwdCheck, "_", "/", -1)

	return (pwdCheck == strings.Split(sentPwd, "{SHA}")[1])
}

func renderUnauthorized(w http.ResponseWriter, error string) {
	w.Header().Set("WWW-Authenticate", "Basic realm=\"\"")
	w.WriteHeader(http.StatusUnauthorized)
	errJson := Error{Message: error}
	WriteIndentedJson(w, errJson, "", "  ")
}
