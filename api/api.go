package api

import(
	"encoding/json"
	"net/http"
)

//  Coin Balance Params
type CoinBalanceParams struct{
	Username string
}

type CoinBalanceResponse struct{
	// success code , usually 200
	Code int
// Account Balance
   Balance int64
}

type Error struct{
	// Error code
	Code int

	// Error message
	message string
}


func writeError(w http.ResponseWriter, message string , code int) {
	resp:= Error{
		Code: code,
		Message : message,
	}

	w.Header().Set("Content-Type","application/json")
    w.WriteHeader(code)

	json.NewEncoder(w).Encode(resp)

}

var (
	RequestEroorHandler = func(w http.ResponseWriter,err error){
		writeError(w,err.Error(),http.StatusBadRequest)
	}
	InternalErrorHandler = func(w http.ResponseWriter){
		writeError(w,"An Unexpected Error Occured.",http.StatusInternalServerError)
	}
)