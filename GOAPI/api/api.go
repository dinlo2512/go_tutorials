package api
import (
    "encoding/json"
    "net/http"
)

// Dữ liệu biến Coins
type CoinBalanceParams struct {
    Username string
}

//Dữ liệu trả về
type CoinBalanceResponse struct {
    //http Code
    Code int
    //Kết quả trả về
    Balance int64
}

//Lỗi trả về
type Error struct {
    //http Code lỗi
    Code int
    //Tin nhắn thông báo
    Message string
}

func writeError (w http.ResponeWriter, message string, code int) {
    var response = Error {
        Code: code,
        Message: message
    }

    w.Header().Set("Content-type", "application/json")
    w.WriteHeader(code)

    json.NewEncoder(w).Encode(response)
}

var (
    RequetErrorHandler = func(w http.ResponeWriter, err error) {
        writeError(w, err.Error(), http.StatusBadRequest)
    }
    InternalErrorHandler = func(w http.ResponeWriter) {
       writeError(w, "Loi server", http.StatusInternalServerError)
    }

)