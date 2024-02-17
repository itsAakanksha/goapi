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
