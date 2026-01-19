package middleware

import (
  "error"
  "net/http"
  "GOAPI/api"
  "GOAPI/internal/tools"
  log "github.com/sirupsen/logrus" //Để dễ debug
)

var UnAuthorization = error.New("Invalid username or token")

func Authorization(next http.Handler) http.Handler {
    return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
        var username = r.URL.Query().Get("username") //Lay tu duong dan
        var token = r.Header.Get("Authorization") //lay tu header api
        var error

        if username == '' || token == '' {
            log.Error(UnAuthorization) //debug
            api.RequetErrorHandler(w, UnAuthorization)
            return
        }

    })
}