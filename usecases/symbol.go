package usecases

import (
	"context"
	"fmt"
	"vdart-test/entity"
)

const ()

//CreateSymbolUC creates symbol redirect instance
func CreateSymbolUC(reader Reader) *SymbolUC {
	return &SymbolUC{
		reader: reader,
	}
}

//Reader to fetch symbol related values
type Reader interface {
	GetCurrency(ctx context.Context) (entity.Currency, error)
	GetSymbol(ctx context.Context) (entity.Symbol, error)
	GetTicker(ctx context.Context) (entity.Ticker, error)
}

//SymbolUC
type SymbolUC struct {
	reader Reader
}

//Redirect creates a redirection url
func (su *SymbolUC) FetchSymbolValue(ctx context.Context) (entity.SymbolResponse, error) {
	currency, err := su.reader.GetCurrency(ctx)
	if err != nil {
		fmt.Println("error while fetching currency : ", err)
	}
	symbol, err := su.reader.GetSymbol(ctx)
	if err != nil {
		fmt.Println("error while fetching symbols : ", err)
	}
	ticker, err := su.reader.GetTicker(ctx)
	if err != nil {
		fmt.Println("error while fetching tickers : ", err)
	}
	resp := entity.SymbolResponse{
		ID:           currency.ID,
		FullName:     currency.FullName,
		Ask:          ticker.Ask,
		Bid:          ticker.Bid,
		Last:         ticker.Last,
		Open:         ticker.Open,
		Low:          ticker.Low,
		High:         ticker.High,
		FreeCurrency: symbol.FreeCurrency,
	}
	return resp, nil
}
