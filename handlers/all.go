package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"vdart-test/entity"
)

//AllSymbol defines FetchAllSymbolValue signature
type AllSymbol interface {
	FetchAllSymbolValue(ctx context.Context) (entity.AllSymbolResponse, error)
}

//AllSymbHandler handles all symbol requests
type AllSymbHandler struct {
	uc AllSymbol
}

//CreateHandler creates authorization handler
func CreateAllHandler(uc AllSymbol) *AllSymbHandler {
	return &AllSymbHandler{
		uc: uc,
	}
}

func (ah *AllSymbHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	respone, err := ah.uc.FetchAllSymbolValue(ctx)
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
