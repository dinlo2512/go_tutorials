package main
import (
    "fmt"
    "net/http"
    "github.com/go-chi/chi"
    "GOAPI/internal/handlers"
    log "github.com/sirupsen/logrus" //Để dễ debug
)
func main() {
    log.SetReportCaller(true)

    var r *chi.Mux = chi.NewRouter()
    handlers.Handler(r)

    fmt.Println("Start API")

    var error = http.ListenAndServe("localhost:8000", r)
    if error != nil {
        log.Error(error)
    }
}