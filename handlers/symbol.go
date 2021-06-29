package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"vdart-test/config"
	"vdart-test/entity"

	"github.com/gorilla/mux"
)

//Symbol defines FetchSymbolValue signature
type Symbol interface {
	FetchSymbolValue(ctx context.Context) (entity.SymbolResponse, error)
}

//SymbHandler handles symbol requests
type SymbHandler struct {
	uc Symbol
}

//CreateHandler creates Symbol handler
func CreateHandler(uc Symbol) *SymbHandler {
	return &SymbHandler{
		uc: uc,
	}
}

func (sh *SymbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	currency := params["symbol"]
	var symbol string
	for _, cur := range config.ServerConfig.Currencies {
		if cur.Currency == currency {
			symbol = cur.Symbol
			break
		}
	}
	if symbol == "" {
		w.Write([]byte("Not Supported Currency"))
	}
	ctx := context.WithValue(context.Background(), "currency", currency)
	ctx = context.WithValue(ctx, "symbol", symbol)
	respone, err := sh.uc.FetchSymbolValue(ctx)
	if err != nil {
		w.Write([]byte("Error while fetching details" + err.Error()))
		return
	}
	res, err := json.Marshal(respone)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	w.Write(res)
}
